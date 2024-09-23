#!/bin/bash

curl -o documentation.json https://devnet-explorer-api.decimalchain.com/api/documentation.json
./gpretty.py
./swagger2code.py
gofmt -w endpoints.go
gofmt -w verify_endpoints.go
mv endpoints.go ../api
