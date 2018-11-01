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
	m := &message.Wrapper{}
	x := make(map[string]interface{})
	for {
		bs, err := r.ReadSlice('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if err = json.Unmarshal(bs, &x); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if err = m.DeserializeJSON(x); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		for _, b := range m.SerializeBytes() {
			fmt.Printf("%02x", b)
		}
		fmt.Println()
	}
}

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
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
	}
	flag.Parse()
	if len(flag.Args()) != 0 {
		flag.Usage()
		os.Exit(2)
	}
}
