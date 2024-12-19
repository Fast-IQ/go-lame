[![PkgGoDev](https://pkg.go.dev/badge/github.com/Fast-IQ/go-lame)](https://pkg.go.dev/github.com/Fast-IQ/go-lame)
[![Go Report Card](https://goreportcard.com/badge/github.com/Fast-IQ/go-lame)](https://goreportcard.com/report/github.com/Fast-IQ/go-lame)


## What's new

  * more lame library code bound
  * better id3 tag support
  * used in a real project of mine so is going to be developed and maintained more rapidly and carefully (and hopefully)

## Example

```go
package main

import (
	"bufio"
	"os"

    "github.com/Fast-IQ/go-lame"
)

func main() {
	of, err := os.Create("output.mp3")
	if err != nil {
		panic(err)
	}
	defer func() { _ = of.Close()}()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	inf, err := os.Open("input.wav")
	if err != nil {
		panic(err)
	}
	defer func() { _ = inf.Close()}()

	r := bufio.NewReader(inf)
	_, _ = r.WriteTo(enc)
}
```
