#!/bin/bash
if ! command -v golint
then
    echo "ERROR: golint not installed"
    echo "ERROR: please run 'go get -u -v golang.org/x/lint/golint'"
    exit 1
fi

golint -set_exit_status "$(go list ./... | grep -v /model/v1)"