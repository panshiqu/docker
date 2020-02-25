#!/bin/bash
v=`cat version`
docker run -d -p 9090:9090 -v /Users/panshiqu/Documents/docker/log:/app/log panshiqu/deploy:v0.0.$v
