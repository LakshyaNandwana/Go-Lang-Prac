#!/bin/bash
set -e

major="1"
minor="0"
patch="0"
prefix="v"

if [ -z "${major}" ] || [ -z "${minor}" ] || [ -z "${patch}" ]  ; then
   echo "unknown version"
   exit 1
fi

git tag | xargs git tag -d &>/dev/null
git fetch --tags 2>/dev/null

tag_latest=$(git tag -l "${prefix}""${major}"."${minor}"."${patch}"-* 2>/dev/null | sort -V | tail -1)

if [ -z "${tag_latest}" ]; then
   tag_new="${prefix}${major}.${minor}.${patch}-1"
   echo "$tag_new"
   exit 0
fi

build_id_latest=$(echo "${tag_latest}" | awk -F"-" '{print $2}')
build_id="$((${build_id_latest} + 1))"

tag_new="${prefix}${major}.${minor}.${patch}-${build_id}"
echo "${tag_new}"