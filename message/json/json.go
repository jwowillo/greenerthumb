package main

import (
	"bufio"
	"encoding/json"
	"errors"
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

// ErrBytes is returned when a byte-string has odd-length.
var ErrBytes = errors.New("byte-string has odd-length")

func logError(err error) {
	greenerthumb.Error("bytes", err)
}

func main() {
	m := &message.Wrapper{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Bytes()

		if len(in)%2 != 0 {
			logError(ErrBytes)
			continue
		}
		bs := make([]byte, len(in)/2)
		for i := 0; i < len(bs); i++ {
			bs[i] = (toByte(in[i*2]) << 4) | toByte(in[i*2+1])
		}

		if err := m.DeserializeBytes(bs); err != nil {
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

// toByte converts a hex-character to its byte value.
func toByte(x byte) byte {
	switch x {
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		return x - 0x30
	case 'a':
		fallthrough
	case 'b':
		fallthrough
	case 'c':
		fallthrough
	case 'd':
		fallthrough
	case 'e':
		fallthrough
	case 'f':
		return x - (0x66 - 0xf)
	}
	return 0
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
