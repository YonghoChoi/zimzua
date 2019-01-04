@echo off

set serviceName=NSOM-WEB
set installerPath=%~dp0\nsom-windows-web
set nsompath=C:\Nexon\nsom-web
set fileName=nsom-web
set displayName="NSOM Web"

xcopy /E %installerPath%\*.* %nsompath%\
copy %~dp0\nssm.exe %nsompath%\nssm.exe
%nsompath%\nssm.exe install %serviceName% %nsompath%\%fileName%.bat
%nsompath%\nssm.exe set %serviceName% DisplayName "%displayName%"
net start %serviceName%
pause