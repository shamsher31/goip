# goip

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goip)
[![Build Status](https://travis-ci.org/shamsher31/goip.svg)](https://travis-ci.org/shamsher31/goip)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)

Tiny utility to deal with IP

### How to install
```go
go get github.com/shamsher31/goip
```

### How to use

```go
package main

import (
  "fmt"
  "github.com/shamsher31/goip"
  "net/http"
)

func reqHandler(w http.ResponseWriter, req *http.Request) {
  // Returns 1.187.109.235, nil
  fmt.Println(ip.Remote(req))
}

func main() {
  // Returns 10.102.29.20, nil
  fmt.Println(ip.Local())

  // Run simple server to test
  http.HandleFunc("/", reqHandler)
  http.ListenAndServe(":8090", nil)
}
```

### Why
This package will be useful when you write application which require you to deal with local or remote
IP of the requested host.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
