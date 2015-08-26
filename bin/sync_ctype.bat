@echo off

SET OLDPATH=%PATH%

SET GOARCH=386
SET CGO_ENABLED=1
SET GHTS_PATH=%GOPATH%\src\github.com\ghts\ghts
SET ZMQ_SRC_DIR=%GHTS_PATH%\3rd_party\zeromq4-x
SET C_INCLUDE_PATH=%ZMQ_SRC_DIR%\include
SET LIBRARY_PATH=%ZMQ_SRC_DIR%\builds\mingw32
SET GCC=%GHTS_PATH%\3rd_party\ruby_devkit_32
SET PATH=%GHTS_PATH%\bin;%GHTS_PATH%\3rd_party\NH_OpenAPI;%GCC%\bin;%GCC%\mingw\bin;%ZMQ_SRC_DIR%\builds\mingw32;%PATH%

cls
copy %GHTS_PATH%\api\NH\wmca_ctype.orig %GHTS_PATH%\api\NH\wmca_ctype_orig.go
go tool cgo -godefs %GHTS_PATH%\api\NH\wmca_ctype_orig.go > %GHTS_PATH%\api\NH\wmca_ctype_pre.go
sed -e 's/int8/byte/g' %GHTS_PATH%\api\NH\wmca_ctype_pre.go > %GHTS_PATH%\api\NH\wmca_ctype.go
del %GHTS_PATH%\api\NH\wmca_ctype_orig.go
del %GHTS_PATH%\api\NH\wmca_ctype_pre.go

SET PATH=%OLDPATH%