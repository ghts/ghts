@echo off

IF NOT DEFINED GOROOT (
    SET GOROOT=C:\Go
)

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

SET PROJECT_ROOT=%GOPATH%\src\github.com\ghts

REM DEFAULT : 64BIT
IF NOT DEFINED GOARCH (
    SET GOARCH=amd64
)

IF /I "%GOARCH%"=="amd64" (
    SET GCC_PATH=C:\msys64\mingw_64
) ELSE (
    SET GCC_PATH=C:\msys64\mingw_32
)

SET CGO_ENABLED=1
SET PATH=%GOROOT%\bin;%GOPATH%\bin;%GCC_PATH%\bin;%PROJECT_ROOT%\batch_scripts;C:\Program Files\Git\bin;C:\msys64\usr\bin

 