#!/bin/bash

docker kill `id -u -n`-empdb
docker rm `id -u -n`-empdb
