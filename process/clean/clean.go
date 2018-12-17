package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

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

	errorHandler := func(err error) { greenerthumb.Error("clean", err) }
	data := make(map[string]map[string][]Pair)

	err := process.Fields(os.Stdin, makeFieldHandler(data), errorHandler)
	if err != nil {
		errorHandler(err)
		ec |= ReadInput
	}

	for name, datas := range clean(data, limit) {
		for field, data := range datas {
			for _, pair := range data {
				fmt.Println(process.Serialize(
					name,
					pair.Timestamp,
					field,
					pair.Value))
			}
		}
	}

	os.Exit(ec)
}

func makeFieldHandler(data map[string]map[string][]Pair) process.FieldHandler {
	return func(name string, ts int64, field string, value float64) {
		if _, ok := data[name]; !ok {
			data[name] = make(map[string][]Pair)
		}
		data[name][field] = append(data[name][field], Pair{
			Timestamp: ts,
			Value:     value,
		})
	}
}

// limit of standard deviations a value can be away from the average.
var limit int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./clean <standard_deviation_limit>")
		p("")
		p("clean reads all input until STDIN is closed and filters")
		p("instances that are more than a passed number of standard")
		p("deviations away from the mean.")
		p("")
		p("An example is:")
		p("")
		p("    ./clean 1")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 0, \"1\": 1}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 2}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 2, \"1\": 3}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 3, \"1\": 4}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 4, \"1\": 5}")
		p("")
		p("    {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 2}")
		p("    {\"Name\": \"A\", \"Timestamp\": 2, \"1\": 3}")
		p("    {\"Name\": \"A\", \"Timestamp\": 3, \"1\": 4}")
		p("")
		p("Error-codes are used for the following:")
		p("")
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

	limitArg, err := strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
	}
	limit = limitArg
}
