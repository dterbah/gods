#!/bin/bash

test_dirs=$(find . -type f -name '*_test.go' -exec dirname {} \; | sort -u)
for dir in $test_dirs; do
    go test "$dir" -v
done