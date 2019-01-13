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
	greenerthumb.Error("greenerthumb-message-bytes", err)
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

		fmt.Printf("%s\n", greenerthumb.BytesToHex(m.SerializeBytes()))
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
		p("Each byte is written in base-16.")
		p("")
		p("CRC and message errors will be written to STDERR.")
		p("")
		p("An example is:")
		p("")
		p("    ./bytes")
		p("")
		p(`    < {"Header": {"Name": "Soil", "Timestamp": 1, "Sender": "A"}, "Moisture": 0.5}`)
		p("    01000000000000000101413f00000083")
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
