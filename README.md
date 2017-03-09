# GoGFAPI
[![gogfapi documentation on GoDoc.org](https://godoc.org/github.com/gluster/gogfapi/gfapi?status.png)](http://godoc.org/github.com/gluster/gogfapi/gfapi)

A GoGFAPI is Go wrapper around libgfapi, a userspace C-library to access GlusterFS volumes.
GoGAPI provides a Go standard library (`os`) like API to access files on GlusterFS volumes.
More information on the API is available on [godoc.org/github.com/gluster/gofapi/gfapi](https://godoc.org/github.com/gluster/gogfapi/gfapi).

> Note: GoGFAPI uses [cgo](https://golang.org/cmd/cgo/) to bind with libgfapi.

## Using GoGFAPI

First ensure that libgfapi is installed on your system. For Fedora and CentOS (and other EL systems) install the `glusterfs-api` package.

Get GoGFAPI by doing a `go get`.
```
go get -u github.com/gluster/gogfapi/gfapi
```

Import `github.com/gluster/gogfapi/gfapi` into your program to use it.

A simple example,
```go
package main

import "github.com/gluster/gogfapi/gfapi"

func main() {
	vol := &gfapi.Volume{}
	if vol.Init("localhost", "testvol") != 0 {
		// handle error
	}
	if vol.Mount() != 0 {
		//handle error
	}
	defer vol.Unmount()

	f, e := vol.Create("testfile")
	if e != nil {
		//handle error
	}
	defer f.Close()

	_, e := f.Write([]byte{"hello"}, 0)
	if e != nil {
		//handle error
	}

	return
}
```
