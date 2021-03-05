#!/bin/bash -ex

for tag in v1 v2 v3 blue green
do
    sed -ie "s|version = \"v1\"|version = \"$tag\"|" main.go
    skaffold build --tag=$tag
    sed -ie "s|version = \"$tag\"|version = \"v1\"|" main.go
done