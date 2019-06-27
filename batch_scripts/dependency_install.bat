@echo off

call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\set_environmental_variables.bat

go get -u github.com/pebbe/zmq3
go get -u github.com/ugorji/go/codec
go get -u golang.org/x/text
