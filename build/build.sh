#!/bin/bash

set -a
. ./.env
set +a

go build
./twitter-bot