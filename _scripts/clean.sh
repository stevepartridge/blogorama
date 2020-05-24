#!/usr/bin/env bash
# set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=`pwd`

imgID=$(docker volume ls | grep blog_ | awk '{print $2}')
if [[ $imgID != "" ]]; then
  docker volume rm $imgID
else
  echo "Nothing to clean"
fi

cd $CUR_DIR