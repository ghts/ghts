@echo off

SET GOARCH=386
SET cgo_enabled=1
SET CC_FOR_TARGET=

SET OLDPATH=%PATH%

SET GCC=D:\DevTools\ruby_devkit_32\
SET PATH=%GCC%bin;%GCC%mingw\bin;%PATH%

SET ZMQ_SRC_DIR=D:\Work\zmq\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32

REM Bootstrapping cgo cross compile
go tool dist install -v runtime
go install -v -a std

cls
go test

SET PATH=%OLDPATH%