@echo off

SET OLDPATH=%PATH%

SET GOARCH=amd64
SET GCC=D:\DevTools\ruby_devkit_64\
SET PATH=%GCC%bin;%GCC%mingw\bin;%PATH%

SET ZMQ_SRC_DIR=D:\Work\zmq\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32

cls
go test -covermode=count -coverprofile=count.out
go tool cover -func=count.out

SET PATH=%OLDPATH%

