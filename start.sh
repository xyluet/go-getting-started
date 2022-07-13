#!/bin/sh
set -e

export XPORT=${PORT}

exec /bin/go-getting-started
