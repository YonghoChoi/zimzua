@echo off
chcp 65001

REM $(aws ecr get-login --no-include-email --region ap-southeast-1)
set APP_NAME=zimzua-api
set BIN_DIR=.\bin\%APP_NAME%
set /p ZIMZUA_VERSION=<..\..\..\version.txt
set IMAGE_NAME=444926633043.dkr.ecr.ap-southeast-1.amazonaws.com/test-api

if exist "%BIN_DIR%\*" (
    rd /s /q  "%BIN_DIR%"
    if errorlevel 1 (
        pause
        goto :EOF
    )
)

md %BIN_DIR%
xcopy /E ..\..\..\..\..\bin\%APP_NAME%\*.* %BIN_DIR%
docker build -t %IMAGE_NAME%:%ZIMZUA_VERSION% .
docker push %IMAGE_NAME%:%ZIMZUA_VERSION%

docker tag %IMAGE_NAME%:%ZIMZUA_VERSION% %IMAGE_NAME%:latest
docker push %IMAGE_NAME%:latest

rd /s /q  "%BIN_DIR%"