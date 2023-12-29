#!/bin/bash

killall go-inertia-app
CGO_ENABLED=0 go build -o httpd/go-inertia-app

httpd/go-inertia-app &
