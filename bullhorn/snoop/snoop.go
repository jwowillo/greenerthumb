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
	// Resolve is the error-code for failing to resolve a UDP-address.
	Resolve = 1 << iota
	// Listen is the error-code for failing to listen for UDP messages.
	Listen = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("snoop", l, args...)
}

func logError(err error) {
	greenerthumb.Error("snoop", err)
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		logError(err)
		os.Exit(Resolve)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logError(err)
		os.Exit(Listen)
	}

	logInfo("receiving broadcasts from %d", port)

	defer conn.Close()
	io.Copy(os.Stdout, conn)
}

var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./snoop <port>")
		p("")
		p("snoop messages from yellers.")
		p("")
		p("An example is:")
		p("")
		p("    ./snoop 5050")
		p("")
		p("Error-codes are used for the following.")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to resolve a UDP address.",
			Resolve))
		p(fmt.Sprintf(
			"    %d = Failed to listen for UDP messages.",
			Listen))
		p("")
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
