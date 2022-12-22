#!/bin/bash

STATE=$1
if [ -z $STATE ] ; then echo run/test/duh; exit; fi
if [ $STATE == "run" ] ; then
go run ./cmd/web/

fi
if [ $STATE == "test" ] ; then

go test -v ./cmd/web/
go test -failfast -v ./cmd/web
go test ./...
go test -v -run="^TestPing$" ./cmd/web/
go test -v -run="^TestHumanDate$/^UTC|CET$" ./cmd/web/
go test -parallel 4 ./...
go test -race ./cmd/web/
fi

