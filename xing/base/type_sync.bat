@echo off

ECHO 1

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

ECHO 2

REM *********** 
REM *  32Bit  *
REM ***********

call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\32.bat

ECHO 3

REM cls
cd %GOPATH%\src\github.com\ghts\ghts\xing\base

ECHO 4

copy type_c.orig type_1.go >NUL

ECHO 5

"%GOROOT%\bin\go.exe" tool cgo -godefs type_1.go > type_2.go

ECHO 6

sed -e 's/uint8/byte/g' type_2.go > type_3.go

ECHO 7

sed -e 's/int8/byte/g' type_3.go > type_c.go

ECHO 8

del type_1.go
del type_2.go
del type_3.go