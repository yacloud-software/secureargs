.PHONY:	client server
all:	client server
client:
	cd client && go install secureargs-client.go `ls -1 *.go|grep -v secureargs-client.go`
server:
	cd server && go install secureargs-server.go `ls -1 *.go|grep -v secureargs-server.go`
