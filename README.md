# Kerio Control API
[![Go Reference](https://pkg.go.dev/badge/github.com/igiant/control.svg)](https://pkg.go.dev/github.com/igiant/control)
## Overview
Client for [Kerio API Control (JSON-RPC 2.0)](https://manuals.gfi.com/en/kerio/api/control/reference/index.html)

Implemented all Administration API for Kerio Control methods

## Installation
```go
go get github.com/igiant/control
```

## Example
```go
package main

import (
	"fmt"
	"log"

	"github.com/igiant/control"
)

func main() {
	config := control.NewConfig("server_addr")
	conn, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	app := &control.ApiApplication{
		Name:    "MyApp",
		Vendor:  "Me",
		Version: "v0.0.1",
	}
	err = conn.Login("user_name", "user_password", app)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = conn.Logout()
		if err != nil {
			log.Println(err)
		}
	}()
	info, err := conn.ProductInfoGet()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"ProductVersion: %s\nOs: %s\n",
		info.VersionString,
		info.OsDescription,
	)
}
```
## Documentation
* [GoDoc](http://godoc.org/github.com/igiant/control)

## RoadMap
* Add tests and search errors
