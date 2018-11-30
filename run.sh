#!/bin/sh
rm block
rm *.db
rm *.db.lock
go build -o block *.go
ls


