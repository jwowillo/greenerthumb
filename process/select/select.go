package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/process"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to parse input.
	ReadInput = 1 << iota
)

func main() {
	var ec int

	includes := make(map[string]interface{})
	for _, include := range includeList {
		includes[include] = struct{}{}
	}

	fieldHandler := makeFieldHandler(includes)
	errorHandler := func(err error) { greenerthumb.Error("select", err) }

	err := process.Fields(os.Stdin, fieldHandler, errorHandler)
	if err != nil {
		errorHandler(err)
		ec |= ReadInput
	}

	os.Exit(ec)
}

func makeFieldHandler(includes map[string]interface{}) process.FieldHandler {
	return func(name string, ts int64, field string, value float64) {
		if _, ok := includes[name]; !ok {
			return
		}

		fmt.Println(process.Serialize(name, ts, field, value))
	}
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

var includeList ListFlags

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./select [--include <include>...]")
		p("")
		p("select messages from STDIN with names in an included set.")
		p("")
		p("An example is:")
		p("")
		p(`    ./select --include "A" --include "B"`)
		p("")
		p(`    < {"Name": "A", "Timestamp": 0, "1": 1}`)
		p(`    {"Name": "A", "Timestamp": 0, "1": 1}`)
		p(`    < {"Name": "B", "Timestamp": 0, "1": 1}`)
		p(`    {"Name": "B", "Timestamp": 0, "1": 1}`)
		p(`    < {"Name": "C", "Timestamp": 0, "1": 1}`)
		p("")
		p("Error-codes are used for the following:")
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")

		os.Exit(2)
	}
	flag.Var(&includeList, "include", "name to include")
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
