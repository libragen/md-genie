#!/bin/bash

app="md-genie"
pkill $app
rm -rf #app
go build
nohup ./$app >> /md_gennie.log 2>&1 &

