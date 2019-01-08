package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// Dial is the error-code for failing to dial a UDP-connections.
	Dial = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("bullhorn-broadcast-server", l, args...)
}

func logError(err error) {
	greenerthumb.Error("bullhorn-broadcast-server", err)
}

func main() {
	conn, err := net.Dial("udp", fmt.Sprintf("255.255.255.255:%d", port))
	if err != nil {
		logError(err)
		os.Exit(Dial)
	}
	defer conn.Close()

	logInfo("broadcasting on port %d", port)

	io.Copy(conn, os.Stdin)
}

var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./server <port>")
		p("")
		p("server broadcasts messages to all clients.")
		p("")
		p("An example that broadcasts 'a' and 'b ' is:")
		p("")
		p("    ./server 5050")
		p("")
		p("    < a")
		p("    < b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to dial a UDP-connection.",
			Dial))
		p("")

		os.Exit(2)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
	}

	portArg, err := strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
	}
	port = portArg
}
