#!/usr/bin/env sh
set -e
protoc types.proto --go_out=build/go --python_out=build/python
