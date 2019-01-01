package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/message"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logError(err error) {
	greenerthumb.Error("bytes", err)
}

func main() {
	m := &message.Wrapper{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Bytes()

		if err := m.DeserializeBytes(in); err != nil {
			logError(err)
			continue
		}
		out, err := json.Marshal(m.SerializeJSON())
		if err != nil {
			logError(err)
			continue
		}

		fmt.Println(string(out))
	}

	if err := scanner.Err(); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./json")
		p("")
		p("json converts bytes messages from STDIN to JSON written to")
		p("STDOUT.")
		p("")
		p("Message errors will be written to STDERR.")
		p("")
		p("An example is:")
		p("")
		p("    ./json")
		p("")
		p("    < 0100000000000000003ebd70a410")
		p("    {\"Name\": \"Soil\", \"Timestamp\": 0, \"Moisture\": 0.37}")
		p("")
		p("The example shows the bytes received as a hex-string for")
		p("documentation purposes")
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
