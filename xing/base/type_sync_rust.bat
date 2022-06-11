@echo off

cls

SET LIBCLANG_PATH=C:\Program Files (x86)\Microsoft Visual Studio\2019\BuildTools\VC\Tools\Llvm\x64\bin
SET PATH_ORIG=%PATH%
SET PATH=%PATH%;C:\Program Files\Git\usr\bin

cd %GOPATH%\src\github.com\ghts\ghts\xing\base

"%HOMEPATH%\.cargo\bin\bindgen.exe" type_c.h -o type_c.rs --no-layout-tests

SET PATH=%PATH_ORIG%