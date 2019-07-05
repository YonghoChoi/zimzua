@echo off
chcp 65001

REM $(aws ecr get-login --no-include-email --region ap-southeast-1)
set APP_NAME=zimzua-mongo
set /p ZIMZUA_VERSION=<..\..\..\version.txt
set IMAGE_NAME=992189553983.dkr.ecr.ap-northeast-2.amazonaws.com/%APP_NAME%

docker build -t %IMAGE_NAME%:%ZIMZUA_VERSION% .
docker push %IMAGE_NAME%:%ZIMZUA_VERSION%

docker tag %IMAGE_NAME%:%ZIMZUA_VERSION% %IMAGE_NAME%:latest
docker push %IMAGE_NAME%:latest