# getopt library For Go

This is a very simple library for [getopt][]-style argument/option
parsing, written in and for [Go][]. It supports both traditional "short"
options (like `-a`, `-b 1`, `-dfe`, etc), and the GNU-style "long"
options (like `--help`, `--fix=everything`, etc).

[I strongly believe](https://www.rollc.at/posts/2023-07-16-getopt/) that
exposing a common and familiar user interface in your programs is
important; the `getopt`-style command line argument parsing is the
single most universally accepted convention, dating back to at least
1980, and widely supported by many platforms, languages, and utilities.

Unfortunately, Go's standard [flag][] module ignores that convention,
and proposes its own. This package offers a simple alternative.

[getopt]: https://en.wikipedia.org/wiki/getopt
[Go]: https://go.dev/
[flag]: https://pkg.go.dev/flag

## Example

You can find more examples in the [examples](/examples) directory of the
source distribution.

```go
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
```

## Documentation

On [pkg.go.dev](https://pkg.go.dev/github.com/rollcat/getopt).

You can also use [godocs](http://godocs.io/github.com/rollcat/getopt),
or the command line:

```shell
go doc github.com/rollcat/getopt
```

## Author and license

[Original code](https://github.com/timtadh/getopt) by Tim Henderson
<<tim.tadh@gmail.com>>.

This fork, and all of its opinionated tweaks, by Kamil Cholewiński
<<kamil@rollc.at>>.

License is [BSD](/LICENSE).
