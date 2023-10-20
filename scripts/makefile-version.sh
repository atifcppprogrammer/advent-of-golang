#!/bin/bash
go_version_new=$1
go_version_old=$(cat go.work | head -n1 | cut -d ' ' -f2)
go work edit -go $go_version_new
git add go.work

dev_setup_file=.github/DEVELOPMENT_SETUP.md
sed -i -e "s/$go_version_old/$go_version_new/g" $dev_setup_file
git add $dev_setup_file

workflow_file=.github/workflows/solution-tests.yml
sed -i -e "s/$go_version_old/$go_version_new/g" $workflow_file
git add $workflow_file

root_path=$(pwd)
module_paths=($(find . -name go.mod | sort | xargs dirname))
for module_path in ${module_paths[@]}; do
  cd $module_path
  go mod edit -go $go_version_new && git add go.mod
  cd $root_path
done

git commit -m "chore: bumped golang version to \`$go_version_new\`"
