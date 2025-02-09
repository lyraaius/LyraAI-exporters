#!/bin/bash
set -e

mkdir -p output/bin output/conf
cp script/bootstrap.sh output/ 2>/dev/null
chmod +x output/bootstrap.sh
go build -o output/bin/app
