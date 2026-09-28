@echo off

REM 이 스크립트가 위치한 폴더로 디렉토리를 변경 (GOPATH 경로 의존 제거)
cd /d "%~dp0"

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

REM *********** 
REM *  32Bit  *
REM ***********

call "%~dp0..\..\batch_scripts\32.bat"

copy type_c.orig type_1.go >NUL
"%GOROOT%\bin\go.exe" tool cgo -godefs type_1.go > type_2.go
sed -e 's/uint8/byte/g' type_2.go > type_3.go
sed -e 's/int8/byte/g' type_3.go > type_c.go

del type_1.go
del type_2.go
del type_3.go