#!/bin/bash
go vet $(go list ./... | grep -v vendor)
go test $(go list ./... | grep -v vendor)
