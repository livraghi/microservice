#!/usr/bin/env bash

set -e
set -x

TAG_COMMIT=$(git rev-list --tags --max-count=1)
TAGS=$([ -n "$TAG_COMMIT" ] && git tag --contains "$TAG_COMMIT" --sort=-version:refname || echo "undefined")
TAG=$(echo "$TAGS" | head -n 1)

if [ "$TAG_COMMIT" != "$(git rev-parse HEAD)" ]
then
  COMMIT="-$(git rev-parse --short HEAD || echo "undefined")"
else
  COMMIT=""
fi

DIRTY=$([ -z "$(git status -s)" ] || echo "-dirty")

echo -e "Current tag: $TAG$COMMIT$DIRTY"
