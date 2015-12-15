package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var start = flag.Int("start", 0, "character position to start from")
var step = flag.Int("step", 1, "number of characters to step by")
var size = flag.Int("size", 80, "number of characters to display")
var delay = flag.Int("delay", 600, "delay in miliseconds between line outputs")

var input string

type circularStringReader struct {
	input    string
	position int
}

func (c *circularStringReader) read() (rune, bool) {
	wrapped := false

	if c.position < 0 || c.position >= len(c.input) {
		wrapped = true
		c.position = c.position % len(c.input)
	}
	ch := c.input[c.position]
	c.position++
	return rune(ch), wrapped
}

func (c *circularStringReader) readString(start, size int) string {
	var result string

	c.position = start

	for i := 0; i < size; i++ {
		ch, _ := c.read()
		result = result + string(ch)

	}

	return result
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\n  %s [OPTIONS] input string\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	csr := circularStringReader{input: strings.Join(flag.Args(), " "), position: 0}

	for i := *start; true; i++ {
		fmt.Println(csr.readString(i*(*step), *size))
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
}
