@echo off

call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\set_environmental_variables.bat

cls
go test -parallel 100