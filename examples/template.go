package main

import (
	"fmt"
	"os"

	"github.com/rollcat/getopt"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-h]\n", os.Args[0])
}

func help() {
	usage()
	fmt.Fprintf(os.Stderr,
		`CHANGEME: This is a template for a Go commandline program.
Options:
    -h, --help  Show this help and exit
`)
}

func main() {
	args, opts, err := getopt.GetOpt(
		os.Args[1:],
		"h",
		[]string{"help"},
	)
	if err != nil || len(args) != 0 {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		usage()
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt.Opt() {
		case "-h":
			fallthrough
		case "--help":
			help()
			os.Exit(0)
		default:
			panic("unexpected argument")
		}
	}

	// Your program goes here.
}
