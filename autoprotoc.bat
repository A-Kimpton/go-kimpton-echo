if not exist ".\echoservice" mkdir .\echoservice
protoc --go_out=plugins=grpc:echoservice ./protos/echoservice.proto