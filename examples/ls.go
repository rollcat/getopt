// This tiny example program serves as a little stress-test,
// attempting to parse (only parse!) all arguments supported by GNU
// coreutils ls(1).
//
// It dumps the Go representation of the parsed arguments to show how
// they possibly could have been interpreted by the real program.
package main

import (
	"fmt"
	"os"

	"github.com/rollcat/getopt"
)

func main() {
	_, opts, err := getopt.GetOpt(
		os.Args[1:],
		"aAbBcCdDfF:gGhHI:klLmnNopqQrRsStT:uUvw:xXZ1",
		[]string{
			"all",        // -a
			"almost-all", // -A
			"author",
			"escape", // -b
			"block-size=",
			"ignore-backups", // -B
			// -c
			// -C
			"color=",    // TODO: optional args
			"directory", // -d
			"dired",     // -D
			// -f
			"classify=", // -F // TODO: optional args
			"file-type",
			"format=",
			"full-time",
			// -g
			"group-directories-first",
			"no-group",       // -G
			"human-readable", // -h
			"si",
			"dereference-command-line", // -H
			"dereference-command-line-symlink-to-dir",
			"hide=",
			"hyperlink=", // TODO: optional args
			"indicator-style=",
			"inode",     // -i
			"ignore=",   // -I
			"kibibytes", // -k
			// -l
			"dereference", // -L
			// -m
			"numeric-uid-gid", // -n
			"literal",         // -N
			// -o
			// -p sets --indicator-style=slash
			"hide-control-chars", // -q
			"show-control-chars",
			"quote-name", // -Q
			"quoting-style=",
			"reverse",   // -r
			"recursive", // -R
			"size",      // -s
			// -S
			"sort=",
			"time=",
			"time-style=",
			// -t
			"tabsize=", // -T
			// -u
			// -U
			// -v
			"width=", // -w
			// -x
			// -X
			"context", // -Z
			"zero",
			// -1
			"help",
			"version",
		},
	)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		fmt.Printf("error: %#v\n", err)
	} else {
		fmt.Printf("opts: %#v\n", opts)
	}
}
