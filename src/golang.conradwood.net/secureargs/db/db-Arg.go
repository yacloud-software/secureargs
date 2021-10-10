package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBArg
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence arg_seq;

Main Table:

 CREATE TABLE arg (id integer primary key default nextval('arg_seq'),name varchar(2000) not null,r_value varchar(2000) not null,repositoryid bigint not null);

Alter statements:
ALTER TABLE arg ADD COLUMN name varchar(2000) not null default '';
ALTER TABLE arg ADD COLUMN r_value varchar(2000) not null default '';
ALTER TABLE arg ADD COLUMN repositoryid bigint not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE arg_archive (id integer unique not null,name varchar(2000) not null,r_value varchar(2000) not null,repositoryid bigint not null);
*/

import (
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/sql"
	"golang.org/x/net/context"
)

type DBArg struct {
	DB *sql.DB
}

func NewDBArg(db *sql.DB) *DBArg {
	foo := DBArg{DB: db}
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBArg) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "insert_DBArg", "insert into arg_archive (id,name, r_value, repositoryid) values ($1,$2, $3, $4) ", p.ID, p.Name, p.Value, p.RepositoryID)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBArg) Save(ctx context.Context, p *savepb.Arg) (uint64, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_Save", "insert into arg (name, r_value, repositoryid) values ($1, $2, $3) returning id", p.Name, p.Value, p.RepositoryID)
	if e != nil {
		return 0, e
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, fmt.Errorf("No rows after insert")
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, fmt.Errorf("failed to scan id after insert: %s", e)
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBArg) SaveWithID(ctx context.Context, p *savepb.Arg) error {
	_, e := a.DB.ExecContext(ctx, "insert_DBArg", "insert into arg (id,name, r_value, repositoryid) values ($1,$2, $3, $4) ", p.ID, p.Name, p.Value, p.RepositoryID)
	return e
}

func (a *DBArg) Update(ctx context.Context, p *savepb.Arg) error {
	_, e := a.DB.ExecContext(ctx, "DBArg_Update", "update arg set name=$1, r_value=$2, repositoryid=$3 where id = $4", p.Name, p.Value, p.RepositoryID, p.ID)

	return e
}

// delete by id field
func (a *DBArg) DeleteByID(ctx context.Context, p uint64) error {
	_, e := a.DB.ExecContext(ctx, "deleteDBArg_ByID", "delete from arg where id = $1", p)
	return e
}

// get it by primary id
func (a *DBArg) ByID(ctx context.Context, p uint64) (*savepb.Arg, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_ByID", "select id,name, r_value, repositoryid from arg where id = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByID: error scanning (%s)", e)
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("No Arg with id %d", p)
	}
	if len(l) != 1 {
		return nil, fmt.Errorf("Multiple (%d) Arg with id %d", len(l), p)
	}
	return l[0], nil
}

// get all rows
func (a *DBArg) All(ctx context.Context) ([]*savepb.Arg, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_all", "select id,name, r_value, repositoryid from arg order by id")
	if e != nil {
		return nil, fmt.Errorf("All: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBArg" rows with matching Name
func (a *DBArg) ByName(ctx context.Context, p string) ([]*savepb.Arg, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_ByName", "select id,name, r_value, repositoryid from arg where name = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByName: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByName: error scanning (%s)", e)
	}
	return l, nil
}

// get all "DBArg" rows with matching Value
func (a *DBArg) ByValue(ctx context.Context, p string) ([]*savepb.Arg, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_ByValue", "select id,name, r_value, repositoryid from arg where r_value = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByValue: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByValue: error scanning (%s)", e)
	}
	return l, nil
}

// get all "DBArg" rows with matching RepositoryID
func (a *DBArg) ByRepositoryID(ctx context.Context, p uint64) ([]*savepb.Arg, error) {
	rows, e := a.DB.QueryContext(ctx, "DBArg_ByRepositoryID", "select id,name, r_value, repositoryid from arg where repositoryid = $1", p)
	if e != nil {
		return nil, fmt.Errorf("ByRepositoryID: error querying (%s)", e)
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("ByRepositoryID: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBArg) Tablename() string {
	return "arg"
}

func (a *DBArg) SelectCols() string {
	return "id,name, r_value, repositoryid"
}

func (a *DBArg) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.Arg, error) {
	var res []*savepb.Arg
	for rows.Next() {
		foo := savepb.Arg{}
		err := rows.Scan(&foo.ID, &foo.Name, &foo.Value, &foo.RepositoryID)
		if err != nil {
			return nil, err
		}
		res = append(res, &foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBArg) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence arg_seq;`,
		`CREATE TABLE arg (id integer primary key default nextval('arg_seq'),name varchar(2000) not null,r_value varchar(2000) not null,repositoryid bigint not null);`,
		`CREATE TABLE arg (id integer primary key default nextval('arg_seq'),name varchar(2000) not null,r_value varchar(2000) not null,repositoryid bigint not null);`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_arg_%d", i), c)
		if e != nil {
			return e
		}
	}
	return nil
}
