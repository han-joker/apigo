package main

import (
	"flag"
)

var argH bool

func init() {
	flag.BoolVar(&argH, "h", false, "Show help information.")
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}
	if argH {
		flag.PrintDefaults()
		return
	}
}
