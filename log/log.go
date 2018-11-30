package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to parse input.
	ReadInput = 1 << iota
	// MakeFile is the error-code for failing to make a file.
	MakeFile = 1 << iota
)

func main() {
	// The error logic is defined here instead of in the makeFile function
	// because we want to avoid calling os.Exit anywhere except main.
	mf := func(s int64) *os.File {
		f, err := makeFile(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(MakeFile)
		}
		return f
	}

	var f *os.File
	current := time.Now().Unix()
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		// First file. This is done separately because we don't want to
		// make files until we have data to write to them.
		if f == nil {
			f = mf(current)
		} else {
			next := time.Now().Unix()
			if next-current > duration {
				current = next

				// Close the previous file and make the next
				// one.
				f.Close()
				f = mf(current)
			}

		}

		fmt.Fprintf(f, "%s\n", scanner.Bytes())
	}

	if f != nil {
		f.Close()
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(ReadInput)
	}
}

func makeFile(s int64) (*os.File, error) {
	return os.Create(fmt.Sprintf("log-%d.log", s))
}

var duration int64

func init() {
	const timeInDay = 86400
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("log messages from STDIN to a file.")
		p("")
		p("An optional duration flag can be specified which sets the")
		p("duration a log file is used before being rotated.")
		p("")
		p("An example is:")
		p("")
		p("    echo 'line' | ./log")
		p("")
		p("    cat log-1543537416.log")
		p("")
		p("    line")
		p("")
		p(" Error-codes are used for the following")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p(fmt.Sprintf(
			"    %d = Failed to make a file.",
			MakeFile))
		p("")
	}
	flag.Int64Var(&duration, "duration", timeInDay, "duration before rotating")
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}
}
