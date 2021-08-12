#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/dg main.go

echo -e \\n

### dg
version=`cat VERSION`

mkdir -p dg-$version
rm bin/dg-$version.tar.gz || true
cp -rf bin/dg config.json sv_start.sh dg-$version/
cp Dockerfile  dg-$version/

mkdir -p dg
cp -rf dg-$version dg
cd dg
ln -s dg-$version latest
chmod 777 latest
cd ../

sleep 1s
tar -zcvf bin/dg-$version.tar.gz dg
rm -rf dg-$version
rm -rf dg

scp bin/dg-$version.tar.gz ubuntu@192.168.2.142:~