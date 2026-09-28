@echo off

IF EXIST C:\Go\bin\go.exe (
    SET GOROOT=C:\Go
) ELSE IF EXIST %ProgramFiles%\Go\bin\go.exe (
    SET GOROOT=%ProgramFiles%\Go
)

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

SET BATCH_SCRIPT_ROOT=%~dp0

REM DEFAULT : 64BIT
IF NOT DEFINED GOARCH (
    SET GOARCH=amd64
)

IF /I "%GOARCH%"=="amd64" (
    SET GCC_PATH=C:\msys64\mingw64
) ELSE (
    SET GCC_PATH=C:\msys64\mingw32
)

SET CGO_ENABLED=1
SET PATH=%GOROOT%\bin;%GOPATH%\bin;%GCC_PATH%\bin;%BATCH_SCRIPT_ROOT%;C:\Program Files\Git\bin;C:\msys64\usr\bin

 