@echo off


SET OLDPATH=%PATH%

SET GOARCH=386
SET CGO_ENABLED=1
SET GHTS_PATH=%GOPATH%\src\github.com\ghts\ghts
SET GCC=%GHTS_PATH%\3rd_party\ruby_devkit_32
SET ZMQ_SRC_DIR=%GHTS_PATH%\3rd_party\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32
SET PATH=GHTS_PATH\bin;%GCC%\bin;%GCC%\mingw\bin;%ZMQ_SRC_DIR%\builds\mingw32;%PATH%

cls
go test -covermode=count -coverprofile=count.out
go tool cover -html=count.out

SET PATH=%OLDPATH%