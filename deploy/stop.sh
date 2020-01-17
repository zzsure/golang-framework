#!/bin/sh
current=`date "+%Y%m%d%H%M%S"`
docker logs golang-golang-framework-root > ./log/$current.log
docker-compose down
