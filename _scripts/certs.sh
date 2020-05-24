#!/usr/bin/env bash
set -o errexit
# set -o nounset
set -o pipefail

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)
CERTS_DIR=${BASE_DIR}/certificates/out

CREATE=$1

if [ ! -f $CERTS_DIR/host.local.crt ] || [ ! -f $CERTS_DIR/host.local.key ] || [ ! -f $CERTS_DIR/Local_Development_Corp._LLC._Inc._CA.crt ]; then
  CREATE=yes
fi

if [[ "$CREATE" == "" ]]; then
  exit;
fi

if [ ! -d $BASE_DIR/certificates ]; then
	mkdir -p $BASE_DIR/certificates
fi
cd $BASE_DIR/certificates

ORG="Local Development Corp. LLC. Inc."
ORG_UNIT="Local Development"
COMMON_NAME="host.local"
DOMAIN="*.host.local,localhost,127.0.0.1"
EXPIRES="2 years"
COUNTRY="US"
PROVINCE="California"
LOCALITY="San Diego"

if [[ ! $(which certstrap) ]]; then
	echo ""
	echo "Install certstrap first, then try again"
	echo ""
	echo "  go get -u github.com/square/certstrap"
	echo ""
	exit 1
fi

echo "=== init cert authority ==="
certstrap init --common-name "${ORG} CA" \
  --expires "${EXPIRES}" \
  --organization "${ORG}" \
  --organizational-unit "${ORG_UNIT}" \
  --country "${COUNTRY}" \
  --province "${PROVINCE}" \
  --locality "${LOCALITY}" \
  --passphrase ""
  # --depot-path "$BASE_DIR/certificates"

echo "=== request certs ==="
certstrap request-cert --common-name "${COMMON_NAME}" \
  --ip 127.0.0.1 \
  --domain ${DOMAIN} \
  --organization "${ORG}" \
  --organizational-unit "${ORG_UNIT}" \
  --country "${COUNTRY}" \
  --province "${PROVINCE}" \
  --locality "${LOCALITY}" \
  --passphrase ""

echo "=== sign certs ==="
certstrap sign "${COMMON_NAME}" \
  --CA "${ORG} CA" \
  --expires "${EXPIRES}" \
  --passphrase ""
  # --intermediate
echo "=== done ==="

cd $CUR_DIR