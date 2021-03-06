package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/process"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to parse input.
	ReadInput = 1 << iota
	// MarshalJSON is the error-code for failing to marshal JSON.
	MarshalJSON = 1 << iota
)

func main() {
	var ec int

	errorHandler := func(err error) { greenerthumb.Error("clean", err) }
	data := make(map[string]map[string][]float64)

	err := process.Fields(os.Stdin, makeFieldHandler(data), errorHandler)
	if err != nil {
		errorHandler(err)
		ec |= ReadInput
	}

	bs, err := json.Marshal(calculateSummaries(data))
	if err != nil {
		ec |= MarshalJSON
		os.Exit(ec)
	}

	fmt.Printf("%s\n", bs)

	os.Exit(ec)
}

func makeFieldHandler(data map[string]map[string][]float64) process.FieldHandler {
	return func(name string, _ int64, field string, value float64) {
		if _, ok := data[name]; !ok {
			data[name] = make(map[string][]float64)
		}
		data[name][field] = append(data[name][field], value)
	}
}

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./summarize")
		p("")
		p("summarize reads all input until STDIN is closed and then")
		p("reports a 5-number-summary for each data-type along with")
		p("how many instances of that data-type were included.")
		p("")
		p("An example is:")
		p("")
		p("    ./summarize")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 0, \"1\": 1}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 2}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 2, \"1\": 3}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 3, \"1\": 4}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 4, \"1\": 5}")
		p("")
		p("    {\"A\": {\"1\": {\"N\": 5, \"Minimum\": 1, \"Q1\": 1.5, \"Median\": 3, \"Q1\": 4.5, \"Maximum\": 5}}}")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p(fmt.Sprintf(
			"    %d = Failed to marshal JSON.",
			MarshalJSON))
		p("")

		os.Exit(2)
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
