#!/bin/bash

# ToDo : 커맨드라인에서 mongo 명령어 실행하도록 수정 필요
# 데이터가 수집되기 전에 인덱스가 생성되어야함
db.createUser({ user: "zimzua", pwd: "zimzua", roles: [ "readWrite", "dbAdmin" ] })
db.storage.createIndex({location: '2dsphere'})