#!/bin/bash

echo 'Installing Goose...'
go get -u github.com/pressly/goose/cmd/goose

echo 'Running migrations...'
goose -dir=./db/migrations mysql gouser:gopassword@/go_crud_rest_api?parseTime=true up
