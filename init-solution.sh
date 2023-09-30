#!/bin/bash
year=$1; day=$(printf %02d $2); name=$(echo ${@:3})
puzzle_path=solutions/year-$year/day-$day

mkdir -p $puzzle_path 
cp ./templates/*.mustache $puzzle_path && cd $puzzle_path
mv main.go.mustache main.go && mv main_test.go.mustache main_test.go

sed -i -e "s/{{YEAR}}/$year/g" main.go
sed -i -e "s/{{DAY}}/$2/g" main.go
sed -i -e "s/{{NAME}}/$name/g" main.go

go mod init adventofgolang/$puzzle_path && cd -
go work use ./$puzzle_path