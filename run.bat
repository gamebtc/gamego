taskkill /f /im agent.exe /t
taskkill /f /im gameclient.exe /t
taskkill /f /im grpcserver.exe /t
taskkill /f /im game100.exe /t
taskkill /f /im gamezjh.exe /t

cd agent
go build
start agent

cd ../gameclient
go build
start gameclient

cd ../grpcserver
go build
start grpcserver

cd ../game100
go build
start game100

cd ../gamezjh
go build
start gamezjh