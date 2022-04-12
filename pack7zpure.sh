#! /bin/bash

rm -f ../erassistant_pure.7z
cp -f main.go.exe 法环助手.exe
7za a ../erassistant_pure.7z 法环助手.exe template README.md
