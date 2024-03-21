#!/bin/bash

chmod +x get_new_version.sh
latest_tag=$(bash get_new_version.sh)
echo "latest tag is $latest_tag"
echo "${latest_tag}" > new_version.txt

go build ./...
go test ./... --short -coverprofile=coverage.out -v 2>&1 | tee protocol.txt | go-junit-report -set-exit-code > report.xml