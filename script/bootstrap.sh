#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/app"
exec $CURDIR/bin/app