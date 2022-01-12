.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} secureargs-client.go `ls -1 *.go|grep -v secureargs-client.go`
server:
	cd server && go install ${LDFLAGS} secureargs-server.go `ls -1 *.go|grep -v secureargs-server.go`
