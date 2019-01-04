#!/bin/bash
if [ "$1" != "zimzua-api" ]; then
  echo "invalid argument. please input zimzua-api."
  exit 1
fi

rm -rf $1-config
echo "completed remove config directory"
systemctl stop $1
systemctl disable $1
rm /etc/systemd/system/$1.service
systemctl daemon-reload
systemctl reset-failed

echo "completed remove service"
