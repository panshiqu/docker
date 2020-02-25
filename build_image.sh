#!/bin/bash
v=`cat version`
n=`expr $v + 1`
s=`printf "%03d\n" $n`
echo $s > version
docker build --tag=panshiqu/deploy:v0.0.$s .
