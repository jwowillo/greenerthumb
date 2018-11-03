package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/jwowillo/greenerthumb/message"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		m := &message.Wrapper{}
		in, err := r.ReadSlice('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		} else if len(in)%2 != 1 {
			fmt.Fprintln(os.Stderr, "bad bytes length")
			continue
		}
		bs := make([]byte, len(in)/2)
		for i := 0; i < len(bs); i++ {
			bs[i] = (toByte(in[i*2]) << 4) | toByte(in[i*2+1])
		}
		if err = m.DeserializeBytes(bs); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		out, err := json.Marshal(m.SerializeJSON())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(string(out))
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
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}
}
