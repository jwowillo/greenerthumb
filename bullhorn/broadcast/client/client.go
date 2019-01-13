package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/bullhorn"
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
	greenerthumb.Info("greenerthumb-bullhorn-broadcast-client", l, args...)
}

func logError(err error) {
	greenerthumb.Error("greenerthumb-bullhorn-broadcast-client", err)
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
	defer conn.Close()

	logInfo("receiving broadcasts from %d", port)

	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			break
		}

		bs, err := bullhorn.CompareSum(buff[:n])
		if err != nil {
			logError(err)
			continue
		}

		fmt.Printf("%s\n", greenerthumb.BytesToHex(bs))
	}
}

var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./client <port>")
		p("")
		p("client receives messages from broadcasters.")
		p("")
		p("An example that receives 'a' and 'b' from a broadcaster is:")
		p("")
		p("    ./client 5050")
		p("")
		p("    a")
		p("    b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to resolve a UDP address.",
			Resolve))
		p(fmt.Sprintf(
			"    %d = Failed to listen for UDP messages.",
			Listen))
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
