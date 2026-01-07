#!/bin/bash

if [ -z "$NOSTR_SECRET_KEY" ]; then
  echo "NOSTR_SECRET_KEY not set"
  exit 0
fi

RELAY="ws://127.0.0.1:7777"
CONTENT="🫀 sovereign relay online — $(date -Is)"

nak event \
  --sec "$(echo $NOSTR_SECRET_KEY | sed 's/^hex://')" \
  --content "$CONTENT" \
  "$RELAY"
