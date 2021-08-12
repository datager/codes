#!/usr/bin/env bash

docker build -t dtmain:v1 .

docker run --rm -idt -p 8991:8991 dtmain:v1