#!/bin/bash   

echo 'Downloadling Glide...'
curl https://glide.sh/get | sh
# rm -rf ./vendor

echo 'Updating all dependencies...'
glide update
