@echo off

cls

del test.exe
C:\msys64\mingw_32\bin\g++ -o test test.cpp
g++ -o test test.cpp
test.exe
