// Package getopt implements getopt-style option parsing.
//
// Function GetOpt parses a command line argument list, and returns a
// list of known options, plus the leftover arguments. An option is
// known if it has been specified either in the string of accepted
// option characters (shortopts), or in the array of accepted long
// options (longopts).
//
// The shortopts string may contain the following elements: individual
// characters, and characters followed by a colon ":", to indicate an
// argument is to follow. For example, an option string "x" recognizes
// an option "-x", and an option string "x:" recognizes an option and
// argument "-x argument".
//
// The longopts array specifies one option per element. Similarly to
// how colon works in shortopts, the option may be followed by an
// equals sign "=", to indicate an expected argument. For example,
// "flag" recognizes the option "--flag", while "flag:" recognizes an
// option and an argument "--flag=argument". The longopts array can be
// empty or nil, to signify that no long options will be processed.
//
// The interpretation of options in the argument list may be cancelled
// by the option "--" (double dash), which causes GetOpt to end
// further argument processing and return the results so far.
//
// The recognized options will be returned in an array of OptArg, in
// the order in which they were encountered.
//
// For example:
//
//     args, opts, err := Getopt(
//         []string{
//             "-h", "-v", "-x", "asdf", "-r",
//             "--flag=arg",
//             "--", "-x", "qwe",
//         },
//         "hvx:r",
//         []string{"help", "flag="},
//     )
//
// Will return the args:
//
//    []string{"-x", "qwe"}  // Parsing terminated by "--"
//
// And the options:
//
//    []OptArg{
//        OptArg{Option: "-h", Argument: ""},
//        OptArg{Option: "-v", Argument: ""},
//        OptArg{Option: "-x", Argument: "asdf"},
//        OptArg{Option: "-r", Argument: ""},
//        OptArg{Option: "--flag", Argument: "arg"},
//    }

package getopt

import "errors"
import "fmt"
import "strings"

// OptArg represents a single parsed option (and its argument, if
// applicable), as parsed by GetOpt.
type OptArg struct {
	Option   string
	Argument string
}

// Opt returns the Option from OptArg. It exists to maintain backward
// compatibility with github.com/timtadh/getopt.
func (o OptArg) Opt() string { return o.Option }

// Arg returns the Argument from OptArg. It exists to maintain
// backward compatibility with github.com/timtadh/getopt.
func (o OptArg) Arg() string { return o.Argument }

// GetOpt parses the provided args, according to shortopts and
// longopts; and returns the leftover args, parsed options with their
// arguments, and (if there was one) any encountered parsing error.
//
// See the package documentation for a description of the shortops and
// longopts formats, as well as how the args are interpreted in their
// context.
func GetOpt(
	args []string,
	shortopts string,
	longopts []string,
) (
	leftovers []string,
	optargs []OptArg,
	err error,
) {
	shorts, err := build_shorts(shortopts)
	if err != nil {
		return nil, nil, err
	}
	longs, err := build_longs(longopts)
	if err != nil {
		return nil, nil, err
	}
	leftovers = args
	skip := false
	emitopt := ""
	for i, arg := range args {
		leftovers = leftovers[1:]
		if arg == "--" {
			if skip {
				err := fmt.Errorf("expected an argument for %q got --", emitopt)
				return nil, nil, err
			}
			break
		} else if skip {
			if len(arg) > 0 && arg[0] == '-' {
				msg := fmt.Sprintf("expected an argument for %q got %v", emitopt, arg)
				return nil, nil, errors.New(msg)
			}
			optargs = append(optargs, OptArg{emitopt, arg})
			skip = false
			continue
		}

		if len(arg) >= 2 && arg[0] == '-' && arg[1] != '-' {
			shargs := arg[1:]
			for i, sharg := range shargs {
				sa := "-" + string(sharg)
				if found, opt, hasarg := short(sa, shorts); found {
					if i != len(shargs)-1 && hasarg {
						msg := fmt.Sprintf("'%v' requires an arg", sa)
						return nil, nil, errors.New(msg)
					} else if hasarg {
						skip = true
						emitopt = opt
					} else {
						optargs = append(optargs, OptArg{opt, ""})
					}
				} else {
					msg := fmt.Sprintf("couldn't find '%v'", sa)
					return nil, nil, errors.New(msg)
				}
			}
		} else if found, opt, oarg, hasarg, err := long(arg, longs); found {
			if err != nil {
				return nil, nil, err
			} else if oarg != "" {
				optargs = append(optargs, OptArg{opt, oarg})
			} else if hasarg {
				skip = true
				emitopt = opt
			} else {
				optargs = append(optargs, OptArg{opt, ""})
			}
		} else {
			if len(arg) > 0 && arg[0] == '-' {
				msg := fmt.Sprintf("couldn't find '%v'", arg)
				return nil, nil, errors.New(msg)
			}
			leftovers = args[i:]
			break
		}
	}
	if skip {
		return nil, nil, fmt.Errorf("expected an argument for %q got end of args", emitopt)
	}

	return leftovers, optargs, nil
}

func build_longs(long []string) (map[string]bool, error) {
	longs := make(map[string]bool)
	for _, opt := range long {
		hasarg := false
		if opt[len(opt)-1] == '=' {
			opt = opt[:len(opt)-1]
			hasarg = true
		}
		opt = "--" + opt
		if _, has := longs[opt]; has {
			msg := fmt.Sprintf(
				"Option %v entered more than one in longs", opt)
			return nil, errors.New(msg)
		} else {
			longs[opt] = hasarg
		}
	}
	return longs, nil
}

func build_shorts(short string) (map[string]bool, error) {
	shorts := make(map[string]bool)
	for i, rc := range short {
		c := string(rc)
		if c == ":" {
			continue
		}
		if _, has := shorts["-"+c]; has {
			msg := fmt.Sprintf(
				"Option %v entered more than one in shorts", c)
			return nil, errors.New(msg)
		} else {
			shorts["-"+c] = false
			if i+1 < len(short) {
				nc := string(short[i+1])
				if nc == ":" {
					shorts["-"+c] = true
				}
			}
		}
	}
	return shorts, nil
}

func short(arg string, shorts map[string]bool) (found bool, opt string, hasarg bool) {
	if hasarg, has := shorts[arg]; has {
		return true, arg, hasarg
	}
	return false, "", false
}

func long(arg string, longs map[string]bool) (
	found bool,
	opt, rarg string,
	hasarg bool,
	err error,
) {
	if i := strings.Index(arg, "="); i != -1 {
		opt = arg[:i]
		rarg = arg[i+1:]
	} else {
		opt = arg
		rarg = ""
	}
	if hasarg, has := longs[opt]; has {
		if !hasarg && rarg != "" {
			msg := fmt.Sprintf(
				"Option %v received an arg, %v, and did not expect one", opt, rarg)
			return false, "", "", false, errors.New(msg)
		}
		return true, opt, rarg, hasarg, nil
	}
	return false, "", "", false, nil
}
