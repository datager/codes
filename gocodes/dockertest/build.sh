#!/usr/bin/env bash
rm dtmain

GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dtmain main.go