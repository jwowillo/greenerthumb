package main

import (
	"fmt"
	"io"
	"os"
)

// MakePrograms makes all the programs in the list so that they write to STDOUT
// to stdOut and STDERR to stdErr.
func MakePrograms(cmds []string, stdOut, stdErr io.Writer) []*Program {
	ps := make([]*Program, 0, len(cmds))
	for _, cmd := range cmds {
		p, err := NewProgram(cmd, stdOut, stdErr)
		if err != nil {
			fmt.Fprintf(stdErr, "%v\n", err)
			continue
		}
		ps = append(ps, p)
	}
	return ps
}

// MapOverPrograms maps the provided function over all the Programs.
//
// Programs that the function fails on are filtered from the returned list.
func MapOverPrograms(ps []*Program, op func(*Program) error) []*Program {
	var noErrors []*Program
	for _, p := range ps {
		if err := op(p); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		noErrors = append(noErrors, p)
	}
	return noErrors
}

// Start the passed program.
func Start(p *Program) error {
	return p.Start()
}

// Wait for the passed program to finish.
func Wait(p *Program) error {
	return p.Wait()
}

// Close the passed program.
func Close(p *Program) error {
	return p.Close()
}
