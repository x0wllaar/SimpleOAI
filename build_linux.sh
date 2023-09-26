#!/usr/bin/env bash
set -exuo pipefail

GOOS=linux GOARCH=386 GO386=softfloat go build -o simpleoai .
