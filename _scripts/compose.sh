#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

CUR_DIR=`pwd`

. ./_scripts/env.sh

cd $BASE_DIR/_docker/

if [[ "$1" == "stop" ]]; then
  docker-compose \
    -f docker-compose.traefik.yaml \
    -f docker-compose.db.yaml \
    down

  exit
fi

if [[ "$NETWORK_NAME" == "" ]]; then

  echo "NETWORK_NAME environment variable is empty or not set.  Please add it to your local.env file and try again."
  exit 1;

fi

set +e
networkExists=$(docker network ls | grep $NETWORK_NAME)
set -e

if [[ "${networkExists}" == "" ]]; then
  echo "Network ${NETWORK_NAME} does not exist, creating..."
  docker network create -d bridge --subnet=${SUBNET_IP}.0/24 $NETWORK_NAME
fi

echo $LOCAL_DOMAIN
docker-compose -f docker-compose.traefik.yaml -f docker-compose.db.yaml up $@ -d

echo "Waiting for traefik..."
while [[ "$(curl --cacert $BASE_DIR/certificates/out/Local_Development_Corp._LLC._Inc._CA.crt -s $BASICAUTH -o /dev/null -w ''%{http_code}'' https://traefik.${LOCAL_DOMAIN}/ping)" != "200" ]]; do sleep 2; done

cd $CUR_DIR

