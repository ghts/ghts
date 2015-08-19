@echo off

SET CGO_ENABLED=1
SET GHTS_PATH=%GOPATH%\src\github.com\ghts\ghts
SET GCC=%GHTS_PATH%\3rd_party\ruby_devkit_32
SET ZMQ_SRC_DIR=%GHTS_PATH%\3rd_party\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32
SET PATH=GHTS_PATH\shared\bin;%GCC%\bin;%GCC%\mingw\bin;%ZMQ_SRC_DIR%\builds\mingw32;%PATH%

cd %GOPATH%\src\github.com\ghts\ghts\3rd_party\zeromq4-x\builds\mingw32

del *.o
del *.a
del *.dll

IF EXIST Makefile (
REM Do Nothing
) ELSE (
copy Makefile.mingw32 Makefile
)

make

go get -u github.com/pebbe/zmq4

cd %GOPATH%\src\github.com\ghts\ghts

