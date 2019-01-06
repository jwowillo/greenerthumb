package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/process"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to parse input fields.
	ReadInput = 1 << iota
)

func logError(err error) {
	greenerthumb.Error("process-flatten", err)
}

func main() {
	var ec int

	windows := make(map[string]map[string]window)

	err := process.Fields(os.Stdin, makeFieldHandler(windows), logError)
	if err != nil {
		logError(err)
		ec |= ReadInput
	}

	// Write out the final values at the right side of the windows.
	for _, fields := range windows {
		for field, w := range fields {
			avg := average(slide(w, w.C))

			s, err := process.Serialize(w.Header, field, avg)
			if err != nil {
				logError(err)
				return
			}

			fmt.Println(s)
		}
	}

	os.Exit(ec)
}

func makeFieldHandler(ws map[string]map[string]window) process.FieldHandler {
	return func(header process.Header, field string, value float64) {
		name, err := header.GetString("Name")
		if err != nil {
			logError(err)
			return
		}

		if _, ok := ws[name]; !ok {
			ws[name] = make(map[string]window)
		}
		if _, ok := ws[name][field]; !ok {
			// The first part of the window needs to be special
			// cased.
			ws[name][field] = window{
				Header: header,
				B:      value,
				C:      value,
			}
		} else {
			w := slide(ws[name][field], value)
			avg := average(w)
			oldHeader := w.Header
			w.Header = header
			ws[name][field] = w

			s, err := process.Serialize(oldHeader, field, avg)
			if err != nil {
				logError(err)
				return
			}

			fmt.Println(s)
		}
	}
}

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./flatten")
		p("")
		p("flatten smooths data by keeping a sliding window of 3")
		p("instances of a data-type and replacing it with a weighted")
		p("average of the 3 instances biased towards the middle")
		p("instance.")
		p("")
		p("An example is:")
		p("")
		p("    ./flatten")
		p("")
		p(`    < {"Header": {"Name": "A"}, "1": 1, "2": 7}`)
		p(`    < {"Header": {"Name": "A"}, "1": 2, "2": 3}`)
		p("")
		p(`    {"Header": {Name": "A"}, "1": 1.16667}`)
		p(`    {"Header": {"Name": "A"}, "2": 6.33334}`)
		p("")
		p(`    < {"Header": {"Name": "B"}, "3": 4}`)
		p(`    < {"Header": {"Name": "A"}, "2": 5}`)
		p("")
		p(`    {"Header": {"Name": "B"}, "3": 4}`)
		p(`    {"Header": {"Name": "A"}, "1": 1.83333}`)
		p(`    {"Header": {"Name": "A"}, "2": 4}`)
		p(`    {"Header": {"Name": "A"}, "2": 4.66667}`)
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
	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
