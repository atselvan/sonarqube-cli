#!/bin/bash

cli="../src/com/privatesquare/sonarqube-cli/sonarqube-cli"
projectName="sonarqube-cli-test"
projectKey="com.privatesquare.test:${projectName}"

#User Api test
echo "===================================================================================================="
echo "Testing getUsersList"
${cli} -getUsersList

echo "===================================================================================================="
echo "Testing printUserDetails with insufficient parameters"
${cli} -printUserDetails -printUserDetails
echo "===================================================================================================="
echo "Testing printUserDetails with non existing user"
${cli} -printUserDetails -printUserDetails -userId test
echo "===================================================================================================="
echo "Testing printUserDetails"
${cli} -printUserDetails -printUserDetails -userId admin

echo "===================================================================================================="
echo "Testing createUser with insufficient parameters"
${cli} -createUser -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0001 -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0001 -name "Test User" -userPassword welcome
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com
echo "===================================================================================================="
echo "Testing createUser"
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0002 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0003 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0004 -name "Test User" -email test@test.com -userPassword welcome
echo "===================================================================================================="
echo "Testing createUser : User already exists"
${cli} -createUser -userId TST0001 -name "Test User" -email test@test.com -userPassword welcome
${cli} -createUser -userId TST0004 -name "Test User" -email test@test.com -userPassword welcome

echo "===================================================================================================="
echo "Testing deactivateUser with insufficient parameters"
${cli} -deactivateUser
echo "===================================================================================================="
echo "Testing deactivateUser with non existing user"
${cli} -deactivateUser -userId test
echo "===================================================================================================="
echo "Testing deactivateUser"
${cli} -deactivateUser -userId TST0001
echo "===================================================================================================="
echo "Testing deactivateUser : User does not exist"
${cli} -deactivateUser -userId TST0001

echo "===================================================================================================="
echo "Testing printUserDetails again for a created user"
${cli} -printUserDetails -printUserDetails -userId TST0004

echo "===================================================================================================="
echo "Cleanup"
${cli} -deactivateUser -userId TST0002
${cli} -deactivateUser -userId TST0003
${cli} -deactivateUser -userId TST0004

echo "===================================================================================================="