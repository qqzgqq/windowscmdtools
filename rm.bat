@echo off
set rmdir=d:\dellist
if not exist %rmdir% mkdir %rmdir%
move %1 %rmdir%>nul 2>&1
if not "errorlevel"=="0" move %1 %1.bak>nul 2>&1
move %1.bak %rmdir%>nul 2>&1
::forfiles.exe /p %rmdir% /d -1 /c "cmd /c del @path"
goto:eof
