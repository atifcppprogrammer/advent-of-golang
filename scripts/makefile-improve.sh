#!/bin/bash
year=$1; day=$(printf %02d $2)
puzzle_path=solutions/year-$year/day-$day
template_name=template.mustache

cp ./templates/commits/improve.mustache $template_name

sed -i -e "s/{{YEAR}}/$year/g" $template_name
sed -i -e "s/{{DAY}}/$day/g" $template_name

git commit -t $template_name && rm $template_name