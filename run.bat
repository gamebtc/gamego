taskkill /f /im agent.exe /t
taskkill /f /im gameclient.exe /t
taskkill /f /im grpcserver.exe /t
taskkill /f /im game100.exe /t
taskkill /f /im gamezjh.exe /t

cd agent
go build
start call agent

cd ../gameclient
go build
start call gameclient

cd ../grpcserver
go build
start call grpcserver

cd ../game100
go build
::start call game100

cd ../gamezjh
go build
::start call gamezjh