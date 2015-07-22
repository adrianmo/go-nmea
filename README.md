# go-nmea

This is a NMEA library for the Go programming language (http://golang.org).

## Installing

### Using `go get`

    go get github.com/adrianmo/go-nmea/nmea

After this command *go-nmea* is ready to use. Its source will be in:

    $GOPATH/src/github.com/adrianmo/go-nmea/nmea

## Example

```golang
package main

import (
	"fmt"
	"github.com/adrianmo/go-nmea/nmea"
)

func main() {
	m, err := nmea.Parse("$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70")
	if err == nil {
		fmt.Printf("%+v\n", m)
	}
}
```
