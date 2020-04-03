@echo off

cls

del test.exe
C:\Rtools\mingw_32\bin\g++ -o test test.cpp
test.exe
