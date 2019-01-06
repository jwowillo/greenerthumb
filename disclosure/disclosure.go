package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// HostLength is the error-code for having a host that is too long.
	HostLength = 1 << iota
)

func logError(err error) {
	greenerthumb.Error("disclosure", err)
}

func serialize(host string) string {
	return fmt.Sprintf(`{"Host":"%s"}`, host)
}

func main() {
	if len(host) > 255 {
		logError(greenerthumb.StringError{String: host, Limit: 255})
		os.Exit(HostLength)
	}

	duration := time.Duration(float64(time.Second) / rate)
	for {
		fmt.Println(serialize(host))
		time.Sleep(duration)
	}
}

var host string
var rate float64

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p(`./disclosure <host> ?--rate <rate>`)
		p("")
		p("disclosure prints the disclosure message body with the")
		p("passed values periodically at the given rate in hertz.")
		p("")
		p("The host can't be longer than 255 characters")
		p("")
		p("The default rate is 5 hertz.")
		p("")
		p("An example after one second is:")
		p("")
		p("    ./disclosure :8080 --rate 1")
		p("")
		p(`    {"Host":":8080"}`)
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Host was too long.",
			HostLength))
		p("")

		os.Exit(2)
	}

	flag.Float64Var(
		&rate,
		"rate",
		5,
		"rate to print disclosures at")

	if len(os.Args) < 2 {
		flag.Usage()
	}

	host = os.Args[1]

	flag.CommandLine.Parse(os.Args[2:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
