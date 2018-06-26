#!/bin/bash

cli="../src/com/privatesquare/sonarqube-cli/sonarqube-cli"
projectName="sonarqube-cli-test"
projectKey="com.privatesquare.test:${projectName}"


#User Api test
${cli} -getUsersList

${cli} -printUserDetails -printUserDetails
${cli} -printUserDetails -printUserDetails -userId test
${cli} -printUserDetails -printUserDetails -userId admin

${cli} -createUser -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0001 -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0001 -name "Test User" -userPassword welcome
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0002 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0003 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0004 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0004 -name "Test User" -email test@test.com -userPassword welcome

${cli} -deactivateUser
${cli} -deactivateUser -userId test
${cli} -deactivateUser -userId TST0001
${cli} -deactivateUser -userId TST0001

${cli} -printUserDetails -printUserDetails -userId TST0004