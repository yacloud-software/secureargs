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

 CREATE TABLE arg (id integer primary key default nextval('arg_seq'),name text not null  ,r_value text not null  ,repositoryid bigint not null  ,artefactid bigint not null  );

Alter statements:
ALTER TABLE arg ADD COLUMN IF NOT EXISTS name text not null default '';
ALTER TABLE arg ADD COLUMN IF NOT EXISTS r_value text not null default '';
ALTER TABLE arg ADD COLUMN IF NOT EXISTS repositoryid bigint not null default 0;
ALTER TABLE arg ADD COLUMN IF NOT EXISTS artefactid bigint not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE arg_archive (id integer unique not null,name text not null,r_value text not null,repositoryid bigint not null,artefactid bigint not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/secureargs"
	"golang.conradwood.net/go-easyops/sql"
	"os"
)

var (
	default_def_DBArg *DBArg
)

type DBArg struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBArg() *DBArg {
	if default_def_DBArg != nil {
		return default_def_DBArg
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBArg(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBArg = res
	return res
}
func NewDBArg(db *sql.DB) *DBArg {
	foo := DBArg{DB: db}
	foo.SQLTablename = "arg"
	foo.SQLArchivetablename = "arg_archive"
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
	_, e := a.DB.ExecContext(ctx, "archive_DBArg", "insert into "+a.SQLArchivetablename+" (id,name, r_value, repositoryid, artefactid) values ($1,$2, $3, $4, $5) ", p.ID, p.Name, p.Value, p.RepositoryID, p.ArtefactID)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBArg) Save(ctx context.Context, p *savepb.Arg) (uint64, error) {
	qn := "DBArg_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (name, r_value, repositoryid, artefactid) values ($1, $2, $3, $4) returning id", p.Name, p.Value, p.RepositoryID, p.ArtefactID)
	if e != nil {
		return 0, a.Error(ctx, qn, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, qn, fmt.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, qn, fmt.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBArg) SaveWithID(ctx context.Context, p *savepb.Arg) error {
	qn := "insert_DBArg"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,name, r_value, repositoryid, artefactid) values ($1,$2, $3, $4, $5) ", p.ID, p.Name, p.Value, p.RepositoryID, p.ArtefactID)
	return a.Error(ctx, qn, e)
}

func (a *DBArg) Update(ctx context.Context, p *savepb.Arg) error {
	qn := "DBArg_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set name=$1, r_value=$2, repositoryid=$3, artefactid=$4 where id = $5", p.Name, p.Value, p.RepositoryID, p.ArtefactID, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBArg) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBArg_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBArg) ByID(ctx context.Context, p uint64) (*savepb.Arg, error) {
	qn := "DBArg_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No Arg with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) Arg with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBArg) TryByID(ctx context.Context, p uint64) (*savepb.Arg, error) {
	qn := "DBArg_TryByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, nil
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) Arg with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBArg) All(ctx context.Context) ([]*savepb.Arg, error) {
	qn := "DBArg_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" order by id")
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("All: error querying (%s)", e))
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
	qn := "DBArg_ByName"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where name = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByName: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByName: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBArg) ByLikeName(ctx context.Context, p string) ([]*savepb.Arg, error) {
	qn := "DBArg_ByLikeName"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where name ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByName: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByName: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBArg" rows with matching Value
func (a *DBArg) ByValue(ctx context.Context, p string) ([]*savepb.Arg, error) {
	qn := "DBArg_ByValue"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where r_value = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByValue: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByValue: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBArg) ByLikeValue(ctx context.Context, p string) ([]*savepb.Arg, error) {
	qn := "DBArg_ByLikeValue"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where r_value ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByValue: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByValue: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBArg" rows with matching RepositoryID
func (a *DBArg) ByRepositoryID(ctx context.Context, p uint64) ([]*savepb.Arg, error) {
	qn := "DBArg_ByRepositoryID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where repositoryid = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByRepositoryID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByRepositoryID: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBArg) ByLikeRepositoryID(ctx context.Context, p uint64) ([]*savepb.Arg, error) {
	qn := "DBArg_ByLikeRepositoryID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where repositoryid ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByRepositoryID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByRepositoryID: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBArg" rows with matching ArtefactID
func (a *DBArg) ByArtefactID(ctx context.Context, p uint64) ([]*savepb.Arg, error) {
	qn := "DBArg_ByArtefactID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where artefactid = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByArtefactID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByArtefactID: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBArg) ByLikeArtefactID(ctx context.Context, p uint64) ([]*savepb.Arg, error) {
	qn := "DBArg_ByLikeArtefactID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,name, r_value, repositoryid, artefactid from "+a.SQLTablename+" where artefactid ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByArtefactID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByArtefactID: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBArg) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.Arg, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBArg) Tablename() string {
	return a.SQLTablename
}

func (a *DBArg) SelectCols() string {
	return "id,name, r_value, repositoryid, artefactid"
}
func (a *DBArg) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".name, " + a.SQLTablename + ".r_value, " + a.SQLTablename + ".repositoryid, " + a.SQLTablename + ".artefactid"
}

func (a *DBArg) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.Arg, error) {
	var res []*savepb.Arg
	for rows.Next() {
		foo := savepb.Arg{}
		err := rows.Scan(&foo.ID, &foo.Name, &foo.Value, &foo.RepositoryID, &foo.ArtefactID)
		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
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
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),name text not null ,r_value text not null ,repositoryid bigint not null ,artefactid bigint not null );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),name text not null ,r_value text not null ,repositoryid bigint not null ,artefactid bigint not null );`,
		`ALTER TABLE arg ADD COLUMN IF NOT EXISTS name text not null default '';`,
		`ALTER TABLE arg ADD COLUMN IF NOT EXISTS r_value text not null default '';`,
		`ALTER TABLE arg ADD COLUMN IF NOT EXISTS repositoryid bigint not null default 0;`,
		`ALTER TABLE arg ADD COLUMN IF NOT EXISTS artefactid bigint not null default 0;`,

		`ALTER TABLE arg_archive ADD COLUMN IF NOT EXISTS name text not null default '';`,
		`ALTER TABLE arg_archive ADD COLUMN IF NOT EXISTS r_value text not null default '';`,
		`ALTER TABLE arg_archive ADD COLUMN IF NOT EXISTS repositoryid bigint not null default 0;`,
		`ALTER TABLE arg_archive ADD COLUMN IF NOT EXISTS artefactid bigint not null default 0;`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}

	// these are optional, expected to fail
	csql = []string{
		// Indices:

		// Foreign keys:

	}
	for i, c := range csql {
		a.DB.ExecContextQuiet(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBArg) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}





