@echo off

IF NOT DEFINED GOROOT (
    SET GOROOT=C:\Go
)

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

SET CGO_ENABLED=1
SET GOARCH=386
SET PATH=%GOROOT%\bin;C:\msys64\mingw32\bin;C:\msys64\usr\bin

cd %GOPATH%\src\github.com\ghts\ghts\xing\c32new
go run xing_C32.go
