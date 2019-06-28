#!/bin/bash
mysql -u root -p < init.sql
mysql -u root -p -D zimzua < example_simple_data.sql