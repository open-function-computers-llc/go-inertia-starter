#!/bin/bash

killall go-inertia-app

npx vite build
CGO_ENABLED=0 go build -ldflags="-X main.Version=`git rev-parse HEAD`" -o cmd/go-inertia-app

./cmd/go-inertia-app &
