protoc --gogofaster_out=plugins=grpc:../ game.proto
protoc --gogofaster_out=plugins=grpc:../ msg_id.proto
::protoc --gogofaster_out=plugins=grpc:../ model.proto
protoc --gogofaster_out=plugins=grpc:../ login.proto
protoc --gogofaster_out=plugins=grpc:../folks/  folks.proto
protoc --gogofaster_out=plugins=grpc:../zjh/  zjh.proto
protoc --gogofaster_out=plugins=grpc:../fish/  fish.proto
msgp.bat
::codecgen.bat