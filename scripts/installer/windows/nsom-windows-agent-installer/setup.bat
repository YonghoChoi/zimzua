@echo off

set serviceName=NSOM-AGENT
set installerPath=%~dp0\nsom-windows-agent
set nsompath=C:\Nexon\nsom-agent
set fileName=nsom-agent
set displayName="NSOM Agent"

xcopy /E /Y %installerPath%\*.* %nsompath%\
copy %~dp0\nssm.exe %nsompath%\nssm.exe
%nsompath%\nssm.exe install %serviceName% %nsompath%\%fileName%.bat
%nsompath%\nssm.exe set %serviceName% DisplayName "%displayName%"
net start %serviceName%
pause