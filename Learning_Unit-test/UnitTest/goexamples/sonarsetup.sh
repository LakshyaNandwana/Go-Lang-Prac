#!/bin/bash

set -e

apt-get update
apt-get --no-install-recommends install -y \
            curl \
            unzip

sonarscanner_version=4.7.0.2747-linux
sonarscanner_zip=sonar-scanner-cli-${sonarscanner_version}.zip
curl -LO "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/${sonarscanner_zip}"
unzip ${sonarscanner_zip} -d /progs
rm ${sonarscanner_zip}
ln -s /progs/sonar-scanner-${sonarscanner_version}/bin/sonar-scanner /usr/local/bin/sonar-scanner