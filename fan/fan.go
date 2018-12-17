package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// OpenStdin is the error-code for failing to open STDIN.
	OpenStdin = 1 << iota
	// StartProgram is the error-code for failing to start a program.
	StartProgram = 1 << iota
	// WaitProgram is the error-code for failing to finish a program.
	WaitProgram = 1 << iota
	// CloseStdout is the error-code for failing to close STDOUT.
	CloseStdout = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logError(err error) {
	greenerthumb.Error("fan", err)
}

func main() {
	var ec int

	// Programs that receive input output to STDOUT.
	inPrograms := MakePrograms(ins, os.Stdout, os.Stderr)
	if len(inPrograms) != len(ins) {
		logError(errors.New("failed to parse an in-program"))
		ec |= OpenStdin
	}

	// Programs that write output write to those that receive input.
	inWriters := make([]io.Writer, 0, len(inPrograms))
	for _, in := range inPrograms {
		inWriters = append(inWriters, in)
	}
	inWriter := NewMultiWriter(inWriters)
	outPrograms := MakePrograms(outs, inWriter, os.Stderr)
	if len(outPrograms) != len(outs) {
		logError(errors.New("failed to parse an out-program"))
		ec |= OpenStdin
	}

	outWriters := make([]io.Writer, 0, len(outPrograms))
	for _, out := range outPrograms {
		outWriters = append(outWriters, out)
	}
	outWriter := NewMultiWriter(outWriters)

	// Write STDIN to out-programs.
	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			in := append(scanner.Bytes(), '\n')

			if _, err := outWriter.Write(in); err != nil {
				logError(err)
			}
		}

		if err := scanner.Err(); err != nil {
			logError(err)
			ec |= ReadInput
		}
	}()

	// Start all the input programs first so they can receive all output.
	startedIns := MapOverPrograms(inPrograms, Start)
	if len(startedIns) != len(inPrograms) {
		logError(errors.New("failed to start an in-program"))
		ec |= StartProgram
	}
	// Start all the output programs once the input programs are running.
	startedOuts := MapOverPrograms(outPrograms, Start)
	if len(startedOuts) != len(outPrograms) {
		logError(errors.New("failed to start an out-program"))
		ec |= StartProgram
	}

	// Wait for the programs that write to the input programs to finish
	// first so none of their output is missed.
	waitedOuts := MapOverPrograms(startedOuts, Wait)
	if len(waitedOuts) != len(startedOuts) {
		logError(errors.New("failed to run an out-program"))
		ec |= WaitProgram
	}
	closedOuts := MapOverPrograms(waitedOuts, Close)
	if len(closedOuts) != len(waitedOuts) {
		logError(errors.New("failed to close an out-program"))
		ec |= CloseStdout
	}

	// The input programs can be waited to finish now that all the output is
	// finished. The input programs STDINs have to be closed first so they
	// don't cause the programs to hang.
	closedIns := MapOverPrograms(startedIns, Close)
	if len(closedIns) != len(startedIns) {
		logError(errors.New("failed to close an in-program"))
		ec |= CloseStdout
	}
	waitedIns := MapOverPrograms(closedIns, Wait)
	if len(waitedIns) != len(closedIns) {
		logError(errors.New("failed to run an in-program"))
		ec |= WaitProgram
	}

	os.Exit(ec)
}

// ListFlags is a flag.Value which can receive multiple values.
type ListFlags []string

// Set the flag.Value by appending to the list.
func (fs *ListFlags) Set(v string) error {
	*fs = append(*fs, v)
	return nil
}

func (fs ListFlags) String() string {
	return strings.Join(fs, ",")
}

// outs are output programs.
var outs ListFlags

// ins are input programs.
var ins ListFlags

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./fan [--out <out>...] [--in <in>...]")
		p("")
		p("fan connects STDOUTS from listed out-programs to STDINs of")
		p("listed in-programs.")
		p("")
		p("An example is:")
		p("")
		p("    ./fan --out 'echo a' --out 'echo b' \\")
		p("        --in 'cat' --in 'cat'")
		p("")
		p("    a")
		p("    a")
		p("    b")
		p("    b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to open a program's STDIN.",
			OpenStdin))
		p(fmt.Sprintf(
			"    %d = Failed to start a program.",
			StartProgram))
		p(fmt.Sprintf(
			"    %d = Failed to wait for a program to finish.",
			WaitProgram))
		p(fmt.Sprintf(
			"    %d = Failed to close a program's STDOUT.",
			CloseStdout))
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")

		os.Exit(2)
	}
	flag.Var(&outs, "out", "out-programs")
	flag.Var(&ins, "in", "in-programs")
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
