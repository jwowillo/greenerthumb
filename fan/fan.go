package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
)

func main() {
	var ec int

	// Programs that receive input output to STDOUT.
	inPrograms := MakePrograms(ins, os.Stdout)
	if len(inPrograms) != len(ins) {
		ec |= 1 << OpenStdin
	}

	// Programs that write output write to those that receive input.
	inWriters := make([]io.Writer, 0, len(inPrograms))
	for _, in := range inPrograms {
		inWriters = append(inWriters, in)
	}
	inWriter := io.MultiWriter(inWriters...)
	outPrograms := MakePrograms(outs, inWriter)
	if len(outPrograms) != len(outs) {
		ec |= 1 << OpenStdin
	}

	outWriters := make([]io.Writer, 0, len(outPrograms))
	for _, out := range outPrograms {
		outWriters = append(outWriters, out)
	}
	outWriter := io.MultiWriter(outWriters...)

	// Write STDIN to out programs.
	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			in, err := r.ReadSlice('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			outWriter.Write(in)
		}
	}()

	// Start all the input programs first so they can receive all output.
	startedIns := MapOverPrograms(inPrograms, Start)
	if len(startedIns) != len(inPrograms) {
		ec |= StartProgram
	}
	// Start all the output programs once the input programs are running.
	startedOuts := MapOverPrograms(outPrograms, Start)
	if len(startedOuts) != len(outPrograms) {
		ec |= StartProgram
	}

	// Wait for the programs that write to the input programs to finish
	// first so none of their output is missed.
	waitedOuts := MapOverPrograms(startedOuts, Wait)
	if len(waitedOuts) != len(startedOuts) {
		ec |= WaitProgram
	}
	closedOuts := MapOverPrograms(waitedOuts, Close)
	if len(closedOuts) != len(waitedOuts) {
		ec |= CloseStdout
	}

	// The input programs can be waited to finish now that all the output is
	// finished. The input programs STDINs have to be closed first so they
	// don't cause the programs to hang.
	closedIns := MapOverPrograms(startedIns, Close)
	if len(closedIns) != len(startedIns) {
		ec |= CloseStdout
	}
	waitedIns := MapOverPrograms(closedIns, Wait)
	if len(waitedIns) != len(closedIns) {
		ec |= WaitProgram
	}

	os.Stdin.Close()

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
		p("fan connects STDOUTS from listed out-programs to STDINs of")
		p("listed in-programs.")
		p("")
		p("Example:")
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
		p("")

		flag.PrintDefaults()
	}
	flag.Var(&outs, "out", "out-programs")
	flag.Var(&ins, "in", "in-programs")
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}
}
