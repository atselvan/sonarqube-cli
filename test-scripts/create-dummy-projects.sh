#!/bin/bash

cli="../src/com/privatesquare/sonarqube-cli/sonarqube-cli"
projectsCount=2300

for i in `seq 1 ${projectsCount}`
do
    ${cli} -createProject -projectName test-project-${i} -projectKey com.privatesquare.test-project-${i}
done