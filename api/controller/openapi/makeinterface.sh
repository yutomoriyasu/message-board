#!/bin/bash

cd `dirname $0`

PACKAGENAME=api

oapi-codegen --old-config-style -generate "types" -package api ./openapi.yml > ./types.gen.go
oapi-codegen --old-config-style -generate "server" -package api ./openapi.yml > ./server.gen.go
oapi-codegen --old-config-style -generate "spec" -package api ./openapi.yml > ./spec.gen.go
