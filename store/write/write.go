package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/store"
)

const (
	_ = iota
	_
	// Connect is the error-code for failing to connect to a store.
	Connect = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("write", l, args...)
}

func logError(err error) {
	greenerthumb.Error("write", err)
}

func main() {
	s, err := store.NewSQLITEStore(path)
	if err != nil {
		logError(err)
		os.Exit(Connect)
	}
	defer s.Close()

	logInfo("connected to store at %s", path)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := s.Write(scanner.Text()); err != nil {
			logError(err)
		}
	}

	if err := scanner.Err(); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

var path string

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./write <path>")
		p("")
		p("write messages from STDIN into the store at the path.")
		p("")
		p("Messages with the same name update existing messages.")
		p("")
		p("An example is:")
		p("")
		p("    ./write store.db")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 0, \"Value\": 0}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 1, \"Value\": 1}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 2, \"Value\": 2}")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to connect to a store.",
			Connect))
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")

		os.Exit(2)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
	}

	path = args[0]
}
