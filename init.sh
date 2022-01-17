#!/bin/bash
rm -rf .idea
go test -v -run TestMain ./test/unit_test.go
