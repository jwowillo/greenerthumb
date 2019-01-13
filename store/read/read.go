package main

import (
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
	// Read is the error-code for failing to read from a store.
	Read = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("greenerthumb-store-read", l, args...)
}

func logError(err error) {
	greenerthumb.Error("greenerthumb-store-read", err)
}

func main() {
	s, err := store.NewSQLITEStore(path)
	if err != nil {
		logError(err)
		os.Exit(Connect)
	}
	defer s.Close()

	logInfo("connected to store at %s", path)

	msgs, err := s.Read()
	if err != nil {
		logError(err)
		os.Exit(Read)
	}

	for _, msg := range msgs {
		fmt.Println(msg)
	}
}

var path string

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./read <path>")
		p("")
		p("read all messages stored at the path.")
		p("")
		p("An example is:")
		p("")
		p("    ./read store.db")
		p("")
		p(`    {"Header": {"Name": "A"}, "Value": 2}`)
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to connect to a store.",
			Connect))
		p(fmt.Sprintf(
			"    %d = Failed to read from a store.",
			Read))
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
