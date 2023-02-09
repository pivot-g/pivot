#!/bin/bash

for file in $(ls *.go ) ; do 
echo Building plugin ${file}
go build -buildmode=plugin ${file}

done 