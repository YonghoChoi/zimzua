@echo off

chcp 65001>NUL

set PERMISSIONS=yongho1037@nexon.co.kr
set SERVICE_NAME=샘플 서비스
set TITLE=샘플
set REDIS=localhost:6379
set DEVELOPMENT_MODE=false
set SMTP_ADDR=relay-intranet-pri.nexon.co.kr

cmd /k %~dp0\nsom-web.exe

