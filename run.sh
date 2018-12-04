#!/bin/sh
rm block
rm *.db
rm *.db.lock
rm *.dat
go build -o block *.go
