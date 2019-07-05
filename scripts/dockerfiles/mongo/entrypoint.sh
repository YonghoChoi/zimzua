#!/bin/bash
mongoimport --authenticationDatabase admin -u zimzua -p zimzua -d zimzua -c storage --file full-storage.json
