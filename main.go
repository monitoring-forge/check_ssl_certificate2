package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jessevdk/go-flags"
)

var version string
var commit string

const (
	OK = iota
	WARNING
	CRITICAL
	UNKNOWN
)

func main() {
	os.Exit(_main())
}

func _main() int {
	opt := &Opt{}
	psr := flags.NewParser(opt, flags.HelpFlag|flags.PassDoubleDash)
	_, err := psr.Parse()
	if opt.Version {
		if commit == "" {
			commit = "dev"
		}
		fmt.Printf(
			"%s-%s\n%s/%s, %s, %s\n",
			filepath.Base(os.Args[0]),
			version,
			runtime.GOOS,
			runtime.GOARCH,
			runtime.Version(),
			commit)
		return OK
	} else if flags.WroteHelp(err) {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		return OK
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return UNKNOWN
	}

	if opt.TCP4 && opt.TCP6 {
		fmt.Printf("Both tcp4 and tcp6 are specified\n")
		return UNKNOWN
	}

	if opt.VerifySNI && opt.SNI == "" {
		fmt.Printf("--sni is required when use --verify-sni\n")
		return UNKNOWN
	}

	msg, err := opt.Verify()
	if err != nil {
		fmt.Println(err.Error())
		return CRITICAL
	}
	fmt.Println(msg)
	return OK
}
