@echo off

REM ***********
REM *  32Bit  *
REM ***********

cls

REM call %GOPATH%\src\github.com\ghts\ghts\batch_scripts\32.bat
REM SET LIBCLANG_PATH=C:\Program Files (x86)\Microsoft Visual Studio\2019\BuildTools\VC\Tools\Llvm\bin
SET LIBCLANG_PATH=C:\Program Files (x86)\Microsoft Visual Studio\2019\BuildTools\VC\Tools\Llvm\x64\bin

cd %GOPATH%\src\github.com\ghts\ghts\xing\base

"%HOMEPATH%\.cargo\bin\bindgen.exe" type_c.h -o type_c.rs
