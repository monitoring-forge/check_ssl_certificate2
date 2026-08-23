package main

import (
	"fmt"
	"os"

	"github.com/monitoring-forge/flagrun"
)

var version string

func (opt *Opt) verifyOptions() error {
	if opt.TCP4 && opt.TCP6 {
		return fmt.Errorf("both tcp4 and tcp6 are specified")
	}

	if opt.VerifySNI && opt.SNI == "" {
		return fmt.Errorf("--sni is required when use --verify-sni")
	}
	return nil
}

func (opt *Opt) Run(_ []string) (any, int) {
	err := opt.verifyOptions()
	if err != nil {
		// options are invalid, return UNKNOWN status and print the error message to stderr
		return err, flagrun.UNKNOWN
	}

	msg, err := opt.Verify()
	if err != nil {
		// ssl verification failed, return CRITICAL status and print the error message to stdout
		fmt.Println(err.Error())
		return "", flagrun.CRITICAL
	}
	fmt.Println(msg)
	return "", flagrun.OK
}

func main() {
	os.Exit(flagrun.Go(&Opt{}, flagrun.Version(version)))
}
