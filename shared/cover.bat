@echo off
go test -covermode=count -coverprofile=count.out
go tool cover -func=count.out