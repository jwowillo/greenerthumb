package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/greenerthumb/process"
)

const (
	_ = iota
	_
	// ParseFields is the error-code for failing to parse input fields.
	ParseFields = 1 << iota
)

func main() {
	var ec int

	errorHandler := func(err error) { fmt.Fprintln(os.Stderr, err) }
	windows := make(map[string]map[string]window)

	err := process.Fields(os.Stdin, makeFieldHandler(windows), errorHandler)
	if err != nil {
		errorHandler(err)
		ec |= ParseFields
	}

	// Write out the final values at the right side of the windows.
	for name, fields := range windows {
		for field, w := range fields {
			avg := average(slide(w, w.C))

			fmt.Println(serialize(name, w.Timestamp, field, avg))
		}
	}

	os.Exit(ec)
}

func makeFieldHandler(ws map[string]map[string]window) process.FieldHandler {
	return func(name string, timestamp int64, field string, value float64) {
		if _, ok := ws[name]; !ok {
			ws[name] = make(map[string]window)
		}
		if _, ok := ws[name][field]; !ok {
			// The first part of the window needs to be special
			// cased.
			ws[name][field] = window{
				Timestamp: timestamp,
				B:         value,
				C:         value,
			}
		} else {
			w := slide(ws[name][field], value)
			avg := average(w)
			oldTimestamp := w.Timestamp
			w.Timestamp = timestamp
			ws[name][field] = w

			fmt.Println(serialize(name, oldTimestamp, field, avg))
		}
	}
}

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("flatten smooths data by keeping a sliding window of 3")
		p("instances of a data-type and replacing it with a weighted")
		p("average of the 3 instances biased towards the middle")
		p("instance.")
		p("")
		p("Example:")
		p("")
		p("    ./flatten")
		p("")
		p("    < {\"Name\": \"A\", \"Timestamp\": 0, \"1\": 1, \"2\": 7}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 2, \"2\": 3}")
		p("")
		p("    {\"Name\": \"A\", \"Timestamp\": 0, \"1\": 1.16667}")
		p("    {\"Name\": \"A\", \"Timestamp\": 0, \"2\": 6.33334}")
		p("")
		p("    < {\"Name\": \"B\", \"Timestamp\": 0, \"3\": 4}")
		p("    < {\"Name\": \"A\", \"Timestamp\": 2, \"2\": 5}")
		p("")
		p("    {\"Name\": \"B\", \"Timestamp\": 0, \"3\": 4}")
		p("    {\"Name\": \"A\", \"Timestamp\": 1, \"1\": 1.83333}")
		p("    {\"Name\": \"A\", \"Timestamp\": 1, \"2\": 4}")
		p("    {\"Name\": \"A\", \"Timestamp\": 2, \"2\": 4.66667}")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to parse input fields.",
			ParseFields))
		p("")
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}
}
