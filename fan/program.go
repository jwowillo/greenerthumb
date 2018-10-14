package main

import (
	"io"
	"os/exec"
	"strings"
)

// Program is an abstraction that runs a program in a shell, can receive input,
// and writes output to a passed io.Writer.
type Program struct {
	cmd    *exec.Cmd
	writer io.WriteCloser
}

// NewProgram made from the raw shell command that writes to the io.Writer.
//
// Returns an error if the program couldn't be configured correctly.
func NewProgram(raw string, out io.Writer) (*Program, error) {
	parts := strings.Split(raw, " ")
	program := parts[0]
	rest := strings.Join(parts[1:], " ")

	var cmd *exec.Cmd

	if len(rest) == 0 {
		cmd = exec.Command(program)
	} else {
		cmd = exec.Command(program, rest)
	}
	writer, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stdout = out
	return &Program{cmd: cmd, writer: writer}, nil
}

// Start the program.
func (p *Program) Start() error {
	return p.cmd.Start()
}

// Write to the program's STDIN.
func (p *Program) Write(bs []byte) (int, error) {
	return p.writer.Write(bs)
}

// Wait for the program to finish.
func (p *Program) Wait() error {
	return p.cmd.Wait()
}

// Close the program's STDIN.
func (p *Program) Close() error {
	return p.writer.Close()
}
