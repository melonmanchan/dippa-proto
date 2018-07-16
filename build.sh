#!/usr/bin/env sh
set -e
mkdir -p build/python
mkdir -p build/go
protoc types.proto --go_out=build/go --python_out=build/python

cp build/python/types_pb2.py ../dippa-nlp
cp build/go/types.pb.go ../dippa-facerec/src/types
cp build/go/types.pb.go ../dippa-publish/src/types
