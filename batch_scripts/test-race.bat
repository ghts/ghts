@echo off

REM ************************************************
REM * Go Race Dectector is supported ONLY ON 64BIT *
REM ************************************************
call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\set_environmental_variables.bat

cls
go test -race