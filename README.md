# az-metadata-viewer

[![Travis CI](https://travis-ci.org/dcasati/az-metadata-viewer.svg?branch=master)](https://travis-ci.org/dcasati/az-metadata-viewer)               
A metadata viewer for Azure written in Golang

## Usage

```bash
$GOPATH/bin/az-metadata-viewer

Compute
------------------------------------------------------------------------------
        Location: westus2
        Name: openbsd20171030222627vm
        OS Type: Linux
        Platform Fault Domain: 0
        Platform Update Domain: 0
        VMID: 8516ec49-cc3a-4a48-b5be-29f8f5ae06c3
        VM Size: Standard_DS2_v2

Network
------------------------------------------------------------------------------
        Private IP: 10.0.0.4
        Public IP: XXXXXXXXX
        Subnet Address: 10.0.0.0/24
        MAC Addr: 000D3AFD1D88

        Private IP: 10.0.1.4
        Subnet Address: 10.0.1.0/24
        MAC Addr: 000D3AF9434F
```

## Install

```bash 
$ go get -u github.com/dcasati/az-metadata-viewer
```
