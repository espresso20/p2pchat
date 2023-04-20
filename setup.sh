#!/bin/sh

if [ ! -f go.mod ]; then
  go mod init chatserver
fi

if [ ! -f go.sum ]; then
  go mod tidy
fi

go mod tidy

