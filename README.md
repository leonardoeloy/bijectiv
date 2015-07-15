# Bijectiv [![Build Status](https://travis-ci.org/leonardoeloy/bijectiv.png)](https://travis-ci.org/leonardoeloy/bijectiv) [![GoDoc](https://godoc.org/github.com/leonardoeloy/bijectiv?status.svg)](http://godoc.org/github.com/leonardoeloy/bijectiv)

A Bijective Encoding Function implementation, useful for URL shorteners. 

# Installation

```bash
go get github.com/leonardoeloy/bijectiv
```

# Usage

```go
package main

import (
	"github.com/leonardoeloy/bijectiv"
	"fmt"
)

func main() {
	b := bijectiv.New()
	fmt.Printf("Encoded value of 100 is %s\n", b.Encode(100))
	fmt.Printf("Decoded value of 'bM' is %d\n", b.Decode("bM"))

	b = bijectiv.NewAlphabet("MYRESTRICTEDALPHABET")
	fmt.Printf("Encoded value of 100 with custom alphabet is %s\n", b.Encode(100))
}

```
