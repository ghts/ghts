@echo off

SET OLDPATH=%PATH%
SET CGO_ENABLED=1
SET GHTS_DIR=%GOPATH%\src\github.com\ghts\ghts
SET GCC_DIR=%GHTS_DIR%\3rd_party\ruby_devkit_32
SET BUILD_DEP_DIR=%GHTS_DIR%\3rd_party\build_dep

REM ****************************************
SET ZMQ_DIR=%GHTS_DIR%\3rd_party\zeromq3-x
REM ****************************************

SET C_INCLUDE_PATH=%BUILD_DEP_DIR%\include;%ZMQ_DIR%\include
SET LIBRARY_PATH=%BUILD_DEP_DIR%\lib;%ZMQ_DIR%\lib
SET NH_OpenAPI_DIR=%GHTS_DIR%\3rd_party\NH_OpenAPI
SET PATH=GHTS_DIR\bin;%GCC_DIR%\bin;%GCC_DIR%\mingw\bin;%BUILD_DEP_DIR%\bin;%ZMQ_DIR%\bin;%NH_OpenAPI_DIR%;%PATH%


cd %ZMQ_DIR%\builds\mingw32
make -f Makefile.mingw32
del -f *.o

copy %ZMQ_DIR%\builds\mingw32\*.dll %BUILD_DEP_DIR%\bin\
copy %ZMQ_DIR%\builds\mingw32\*.a %BUILD_DEP_DIR%\lib\
copy %ZMQ_DIR%\include\* %BUILD_DEP_DIR%\includ\

go get -u github.com/pebbe/zmq3
REM go get -u github.com/djimenez/iconv-go

SET PATH=%OLDPATH%
