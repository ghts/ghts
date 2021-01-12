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

copy types_c.orig types_1.go >NUL

%GOROOT%\bin\go.exe tool cgo -godefs types_1.go > types_2.go
sed -e 's/uint8/byte/g' types_2.go > types_3.go
sed -e 's/int8/byte/g' types_3.go > types_c.go
del types_1.go
del types_2.go
del types_3.go