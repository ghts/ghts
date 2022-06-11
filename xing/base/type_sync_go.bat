@echo off

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

REM *********** 
REM *  32Bit  *
REM ***********

call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\32.bat

REM cls
cd %GOPATH%\src\github.com\ghts\ghts\xing\base

copy type_c.orig type_1.go >NUL
"%GOROOT%\bin\go.exe" tool cgo -godefs type_1.go > type_2.go
sed -e 's/uint8/byte/g' type_2.go > type_3.go
sed -e 's/int8/byte/g' type_3.go > type_c.go

del type_1.go
del type_2.go
del type_3.go