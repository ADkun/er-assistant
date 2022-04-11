#! /bin/bash

rm -f ../erassistant.zip
mv main.go.exe 法环助手.exe
zip -q -r ../erassistant.zip 法环助手.exe tools/ mods/ bak/config.ini
