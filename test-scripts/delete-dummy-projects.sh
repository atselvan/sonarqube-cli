#!/bin/bash

cli="../src/com/privatesquare/sonarqube-cli/sonarqube-cli"
projectsCount=10

while getopts ":v" opt; do
  case $opt in
    v)
      verbose="-verbose"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

for i in `seq 1 ${projectsCount}`
do
    ${cli} -deleteProject -projectKey com.privatesquare.test-project-${i} ${verbose}
done