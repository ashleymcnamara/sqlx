#!/bin/bash

set -e

docker run --name empdb -e MYSQL_DATABASE=employees -e MYSQL_USER=docker -e MYSQL_PASSWORD=docker -p 3306:3306 -e MYSQL_ROOT_PASSWORD=etrade -d bketelsen/sqlx
