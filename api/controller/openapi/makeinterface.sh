#!/bin/bash

cd `dirname $0`

PACKAGENAME=openapi

oapi-codegen --old-config-style -generate "types" -package ${PACKAGENAME} ./openapi.yml > ./types.gen.go
oapi-codegen --old-config-style -generate "server" -package ${PACKAGENAME} ./openapi.yml > ./server.gen.go
oapi-codegen --old-config-style -generate "spec" -package ${PACKAGENAME} ./openapi.yml > ./spec.gen.go
