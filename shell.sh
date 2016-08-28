#!/bin/bash

docker exec -it empdb sh -c 'exec mysql -udocker -pdocker'
