#!/bin/bash   

echo 'Downloadling dep...'
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# rm -rf ./vendor

echo 'Checking dependencies...'
dep check

echo 'Updating all dependencies...'
dep ensure -update
