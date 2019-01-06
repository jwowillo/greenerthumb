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

func logError(err error) {
	greenerthumb.Error("process-select", err)
}

func main() {
	var ec int

	includes := make(map[string]interface{})
	for _, include := range includeList {
		includes[include] = struct{}{}
	}

	fieldHandler := makeFieldHandler(includes)

	err := process.Fields(os.Stdin, fieldHandler, logError)
	if err != nil {
		logError(err)
		ec |= ReadInput
	}

	os.Exit(ec)
}

func makeFieldHandler(includes map[string]interface{}) process.FieldHandler {
	return func(header process.Header, field string, value float64) {
		name, err := header.GetString("Name")
		if err != nil {
			logError(err)
			return
		}

		if _, ok := includes[name]; !ok {
			return
		}

		s, err := process.Serialize(header, field, value)
		if err != nil {
			logError(err)
			return
		}

		fmt.Println(s)
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
		p(`    < {"Header": {"Name": "A"}, "1": 1}`)
		p(`    {"Header": {"Name": "A"}, "1": 1}`)
		p(`    < {"Header": {"Name": "B"}, "1": 1}`)
		p(`    {"Header": {"Name": "B"}, "1": 1}`)
		p(`    < {"Header": {"Name": "C}", "1": 1}`)
		p("")
		p("Error-codes are used for the following:")
		p("")
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
