@echo off

IF NOT DEFINED GOROOT (
    SET GOROOT=C:\Go
)

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

SET CGO_ENABLED=1
SET GOARCH=386
SET PATH=%GOROOT%\bin;C:\Rtools\bin;C:\Rtools\mingw_32\bin;C:\Program Files\Git\bin

cd %GOPATH%\src\github.com\ghts\ghts\xing\c32
go run xing_C32.go
