#!/bin/bash

set -e

docker run --name `id -u -n `-empdb -e MYSQL_DATABASE=employees -e MYSQL_USER=docker -e MYSQL_PASSWORD=docker -P -e MYSQL_ROOT_PASSWORD=etrade -d bketelsen/sqlx
