@echo off

REM ************************************************
REM *  elevate to admin and also stay in the correct directory *
REM ************************************************
REM https://stackoverflow.com/questions/6811372/how-to-code-a-bat-file-to-always-run-as-admin-mode

set "params=%*"
cd /d "%~dp0" && ( if exist "%temp%\getadmin.vbs" del "%temp%\getadmin.vbs" ) && fsutil dirty query %systemdrive% 1>nul 2>nul || (  echo Set UAC = CreateObject^("Shell.Application"^) : UAC.ShellExecute "cmd.exe", "/k cd ""%~sdp0"" && %~s0 %params%", "", "runas", 1 >> "%temp%\getadmin.vbs" && "%temp%\getadmin.vbs" && exit /B )


REM ************************************************
REM * Activate Windows Time Service *
REM ************************************************
REM w32tm /unregister
REM w32tm /register
sc config w32time start= auto
net stop w32time
net start w32time

REM ************************************************
REM * Actually sync time to NTP server
REM ************************************************
w32tm /resync /nowait