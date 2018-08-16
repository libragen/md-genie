#!/bin/bash

app="md-genie"
pkill $app
rm -rf $app
go build
nohup ./$app >> ./gennie.log 2>&1 &
ps -ef | grep md-genie

