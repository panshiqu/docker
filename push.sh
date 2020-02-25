#!/bin/bash
v=`cat version`
docker push panshiqu/deploy:v0.0.$v
