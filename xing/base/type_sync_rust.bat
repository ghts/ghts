@echo off

cls

REM 32비트에 맞게 하려면.
REM >rustup target add i686-pc-windows-msvc

SET LIBCLANG_PATH=C:\Program Files (x86)\Microsoft Visual Studio\2019\BuildTools\VC\Tools\Llvm\x64\bin
SET PATH_ORIG=%PATH%
SET PATH=%PATH%;C:\Program Files\Git\usr\bin

cd %GOPATH%\src\github.com\ghts\ghts\xing\base

"%HOMEPATH%\.cargo\bin\bindgen.exe"  --no-layout-tests type_c.h -o type_c.rs -- --target=i686-pc-windows-msvc

SET PATH=%PATH_ORIG%