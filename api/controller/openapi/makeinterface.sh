#!/bin/bash

cd `dirname $0`

PACKAGENAME=openapi

oapi-codegen -generate "types" -package ${PACKAGENAME} *.yml > ./types.gen.go
oapi-codegen -generate "server" -package ${PACKAGENAME} *.yml > ./server.gen.go
oapi-codegen -generate "spec" -package ${PACKAGENAME} *.yml > ./spec.gen.go