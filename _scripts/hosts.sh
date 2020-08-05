#!/usr/bin/env bash
# set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

CUR_DIR=`pwd`

HOSTS="
host.local
blog.host.local
traefik.host.local
"

for host in $HOSTS
do

  if grep -Fxq "127.0.0.1 $host" /etc/hosts
  then
    echo "$host"
  else
    echo "Did not find $host"
    echo "  Attempting to add 127.0.0.1 $host to /etc/hosts"
    echo "  (this may ask for your password)"
    sudo -- sh -c "echo 127.0.0.1 $host >> /etc/hosts"
  fi

done

cd $CUR_DIR
