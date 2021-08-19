#!/bin/sh
# for Cleanup
rm -r bin/*

# for Linux
GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -v
mv opengauss_exporter bin/opengauss_exporter

# for Windows
GOOS="windows" GOARCH="amd64" CGO_ENABLED=0 go build -v
mv opengauss_exporter.exe bin/opengauss_exporter.exe

# for MacOS
GOOS="darwin" GOARCH="amd64" CGO_ENABLED=0 go build -v
mv opengauss_exporter bin/opengauss_exporter_mac

