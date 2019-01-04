@echo off

set REDIS_ADDR=localhost:6379
set SVN_BIN=C:\Program Files\TortoiseSVN\bin\svn.exe

cmd /k %~dp0\nsom-agent.exe

