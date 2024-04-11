#!/bin/bash

git pull;
make build;
VERSION=$(git describe --tags --always)
if [ ! -f "version" ]; then
    echo "version file not found. Creating a new one."
    echo "0.0.0" >> version.out
    exit 0
fi
local_version=$(cat version)
if [ "$local_version" == "$VERSION" ]; then
    exit 0
else
    echo "$VERSION" >> version.out
fi

go_id=`ps -ef|grep "./bin/nft-market-backend -conf" |grep -v "grep" | awk '{print $2}'`
if [ -z "$go_id" ];
then
    echo "[go pid not found]"
else
    kill  $go_id
    echo "killed $go_id"
fi
sleep 2
mkdir -p ./log
touch ./log/"$VERSION"_start.log
nohup ./bin/nft-market-backend -conf ./configs >> ./log/"$VERSION"_start.log 2>&1 &
sleep 2
new_id=`ps -ef|grep "./bin/nft-market-backend -conf" |grep -v "grep" | awk '{print $2}'`
if [ -n "$new_id" ];
then
echo "[$new_id bin/nft-market-backend start ok !!!]"
fi