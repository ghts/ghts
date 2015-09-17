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

SET NH_PATH=%GHTS_PATH%\api\nh\internal

cd %NH_PATH%
copy %NH_PATH%\wmca_ctype.orig %NH_PATH%\wmca_ctype_orig.go
go tool cgo -godefs %NH_PATH%\wmca_ctype_orig.go > %NH_PATH%\wmca_ctype.go
go tool cgo -godefs %NH_PATH%\wmca_ctype_orig.go > %NH_PATH%\wmca_ctype_pre.go
sed -e 's/int8/byte/g' %NH_PATH%\wmca_ctype_pre.go > %NH_PATH%\wmca_ctype.go
del %NH_PATH%\wmca_ctype_orig.go
del %NH_PATH%\wmca_ctype_pre.go

SET PATH=%OLDPATH%