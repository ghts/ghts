@echo off

SET ZMQ_SRC_DIR=D:\Work\zmq\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32

cls
go test -race