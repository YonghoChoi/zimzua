#!/bin/sh

export DEVELOPMENT_MODE=true

cd /opt/zimzua
/opt/zimzua/zimzua-api >> /var/log/zimzua-api.log
