package main

import (
	"github.com/rollcat/getopt"
	"os"
)

func main() {
	args, opts, err := getopt.GetOpt(
		os.Args[1:],
		"hv",
		nil,
	)
	if err != nil || len(args) > 0 {
		println("Usage: program [-hv]")
		os.Exit(1)
	}
	for _, opt := range opts {
		switch opt.Opt() {
		case "-v":
			println("Version 0.1")
			os.Exit(0)
		case "-h":
			println("Usage: program [-hv]")
			os.Exit(0)
		default:
			panic("unexpected argument")
		}
	}
}
