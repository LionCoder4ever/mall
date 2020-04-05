#!/usr/bin/env bash
cd `dirname $0`
echo "sql migration start"
go run main.go --conf ../../../cmd/test.toml
echo "sql migration end"
