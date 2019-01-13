package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
	// SenderLength is the error-code for having a sender that is too long.
	SenderLength = 1 << iota
)

func logError(err error) {
	greenerthumb.Error("greenerthumb-message-header", err)
}

func main() {
	if len(sender) > 255 {
		logError(greenerthumb.StringError{String: sender, Limit: 255})
		os.Exit(SenderLength)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var x map[string]interface{}

		bs := scanner.Bytes()

		if err := json.Unmarshal(bs, &x); err != nil {
			logError(err)
			continue
		}

		x["Header"] = map[string]interface{}{
			"Name":      name,
			"Sender":    sender,
			"Timestamp": time.Now().Unix(),
		}

		bs, err := json.Marshal(x)
		if err != nil {
			logError(err)
		}

		fmt.Printf("%s\n", bs)
	}

	if err := scanner.Err(); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

var (
	name   string
	sender string
)

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./header <name> <sender>")
		p("")
		p("header wraps JSON objects with the message header.")
		p("")
		p("The sender can't be longer than 255 characters.")
		p("")
		p("Errors will be written to STDERR.")
		p("")
		p("An example is:")
		p("")
		p("./header name sender")
		p("")
		p(`    < {"Key": "Value"}`)
		p(`    {"Header": {"Name": "name", "Sender": "sender", "Timestamp": 0}, "Key": "Value"}`)
		p("")
		p("Error-coes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Sender was too long.",
			SenderLength))
		p("")

		os.Exit(2)
	}

	if len(os.Args) != 3 {
		flag.Usage()
	}

	name = os.Args[1]
	sender = os.Args[2]
}
