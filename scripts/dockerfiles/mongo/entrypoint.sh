#!/bin/bash
mongo -u zimzua -p zimzua < /opt/init.js
mongoimport --authenticationDatabase admin -u zimzua -p zimzua -d zimzua -c storage --file /opt/full-storage.json
