#!/bin/bash
git_branch=$(git rev-parse --abbrev-ref HEAD)
is_puzzle_branch=$(echo $git_branch | grep -oE '(solution|improvement)\/[0-9]{4}\/[0-9]{2}' | wc -l)

if [[ $is_puzzle_branch -eq 1 ]]; then 
  year=$(echo $git_branch | cut -d '/' -f2)
  day0=$(echo $git_branch | cut -d '/' -f3)
  test_paths=(solutions/year-$year/day-$day0)
else 
  test_paths=($(find . -name go.mod | sort | xargs dirname))
fi

root_path=$(pwd)
for test_path in ${test_paths[@]}; do
  cd $test_path && go test -v ./... 
  if [[ $? -ne 0 ]]; then exit 1; fi
  cd $root_path
done