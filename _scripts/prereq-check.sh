#!/bin/bash

REQUIRED_GO_VERSION=1.13
REQUIRED_DOCKER_VERSION=18.06
REQUIRED_COMPOSE_VERSION=1.25

check_version()
{
    local version=$1 check=$2
    local winner=$(echo -e "$version\n$check" | sed '/^$/d' | sort -t. -s -k 1,1nr -k 2,2nr -k 3,3nr -k 4,4nr | head -1)
    [[ "$winner" = "$version" ]] && return 0
    return 1
}

check_prereq()
{
    if [ ! $# == 3 ]; then
        echo "Unable to determine '$1' version! Ensure that '$1' is in your path and try again." && exit 1
    fi

    local dependency=$1 required_version=$2 installed_version=$3

    type $dependency >/dev/null 2>&1 || { echo >&2 "Requires '$dependency' but it doesn't appear to be installed.  Aborting."; exit 1; }

    if check_version $installed_version $required_version; then
        echo "$dependency minimum requirement met. Required: $required_version, Found: $installed_version"
    else
        echo "WARNING! Did not find the minimum supported version of '$dependency' installed. Required: $required_version, Found: $installed_version"
        echo "We highly recommend stopping installation and updating dependencies before continuing"
        read -p "Enter Y to continue anyway (not recommended)." -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]
        then
            exit 1
        fi
    fi
}

echo "Checking prerequisites"

GO_VERSION=$(sed -ne 's/[^0-9]*\(\([0-9]\.\)\{0,4\}[0-9][^.]\).*/\1/p' <<< $(go version))
DOCKER_VERSION=$(docker version --format '{{.Server.Version}}' | sed 's/[a-z-]//g')
COMPOSE_VERSION=$(sed -ne 's/[^0-9]*\(\([0-9]\.\)\{0,4\}[0-9][^.]\).*/\1/p' <<< $(docker-compose -v))

check_prereq 'go' $REQUIRED_GO_VERSION $GO_VERSION
check_prereq 'docker' $REQUIRED_DOCKER_VERSION $DOCKER_VERSION
check_prereq 'docker-compose' $REQUIRED_COMPOSE_VERSION $COMPOSE_VERSION
