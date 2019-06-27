@echo off

call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\set_environmental_variables.bat

cls
go test -covermode=count -coverprofile=count.out
go tool cover -html=count.out