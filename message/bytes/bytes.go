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
		x := make(map[string]interface{})
		bs := scanner.Bytes()

		if err := json.Unmarshal(bs, &x); err != nil {
			logError(err)
			continue
		}
		if err := m.DeserializeJSON(x); err != nil {
			logError(err)
			continue
		}

		fmt.Printf("%s\n", m.SerializeBytes())
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
		p("./bytes")
		p("")
		p("bytes converts JSON messages from STDIN to bytes written to")
		p("STDOUT.")
		p("")
		p("CRC and message errors will be written to STDERR.")
		p("")
		p("An example is:")
		p("")
		p("    ./bytes")
		p("")
		p("    < {\"Name\": \"Soil\", \"Timestamp\": 0, \"Moisture\": 0.37}")
		p("    0100000000000000003ebd70a410")
		p("")
		p("The example shows the bytes written as a hex-string for")
		p("documentation purposes.")
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
