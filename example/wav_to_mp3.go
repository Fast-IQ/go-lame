package main

import (
	"bufio"
	"os"

	"github.com/Fast-IQ/go-lame"
)

func main() {
	of, err := os.Create("example/output.mp3")
	if err != nil {
		panic(err)
	}
	defer func() { _ = of.Close() }()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	inf, err := os.Open("example/input.wav")
	if err != nil {
		panic(err)
	}
	defer func() { _ = inf.Close() }()

	r := bufio.NewReader(inf)
	_, _ = r.WriteTo(enc)
}
