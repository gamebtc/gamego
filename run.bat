cd agent
go build
start agent

cd ../gameclient
go build
start gameclient

cd ../grpcserver
go build
start grpcserver

cd ../game
go build
start game