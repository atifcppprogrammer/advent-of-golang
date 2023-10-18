#!/bin/bash
root_path=$(pwd)
test_paths=($(find . -name go.mod | sort | xargs dirname))
for test_path in ${test_paths[@]}; do
  cd $test_path && go test -v ./... 
  if [[ $? -ne 0 ]]; then exit 1; fi
  cd $root_path
done