#! /bin/bash

rm -f ../erassistant.7z
mv main.go.exe 法环助手.exe
7za a ../erassistant.7z 法环助手.exe tools/ mods/ bak/config.ini
