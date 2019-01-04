@echo off

set serviceName=NSOM-AGENT
set dirName=nsom-agent

C:\Nexon\%dirName%\nssm.exe stop %serviceName%
C:\Nexon\%dirName%\nssm.exe remove %serviceName% confirm
@RD /S /Q "C:\Nexon\%dirName%"
REM del /s C:\Nexon\%dirName%nssm.exe
pause