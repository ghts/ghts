@echo off

REM *******************************
REM * Bootstrap 386 Cross Compile *
REM *******************************
call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\32.bat

SET OLDPATH=%PATH%
SET PATH=%GCC_PATH%\bin;%GCC_PATH%\mingw\bin;%DLL_PATH%;%PATH%

go tool dist install -v runtime
go install -v -a std

SET PATH=%OLDPATH%

REM *********************************
REM * Bootstrap amd64 Cross Compile *
REM *********************************
call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\64.bat

SET OLDPATH=%PATH%
SET PATH=%GCC_PATH%\bin;%GCC_PATH%\mingw\bin;%DLL_PATH%;%PATH%

go tool dist install -v runtime
go install -v -a std

SET PATH=%OLDPATH%