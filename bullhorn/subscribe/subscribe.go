package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

const (
	_ = iota
	_
	// Resolve is the error-code for failing to resolve a UDP-address.
	Resolve = 1 << iota
	// Listen is the error-code for failing to listen on a UDP-address.
	Listen = 1 << iota
	// Dial is the error-code for failing to dial a TCP-address.
	Dial = 1 << iota
)

func main() {
	handler := func(err error) { fmt.Fprintln(os.Stderr, err) }
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		handler(err)
		os.Exit(Resolve)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		handler(err)
		os.Exit(Listen)
	}

	tcpConn, err := net.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", publishHost, publishPort))
	if err != nil {
		handler(err)
		os.Exit(Dial)
	}

	fmt.Fprintf(tcpConn, fmt.Sprintf("%d\n", port))

	go func() {
		ioutil.ReadAll(tcpConn)
		conn.Close()
	}()

	io.Copy(os.Stdout, conn)
}

// port to receive data on.
var port int

// publishHost of the bullhorn publisher.
var publishHost string

// publishPort of the bullhorn publisher.
var publishPort int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("subscribe to a bullhorn publisher on a network and write")
		p("its data to STDOUT as a bullhorn subscriber.")
		p("")
		p("A port to receive data, the publisher's IP-address, and the")
		p("publisher's port must be passed.")
		p("")
		p("An example that receives 'a' and 'b' from a publisher is:")
		p("")
		p("    ./subscribe 8080 127.0.0.1 5050")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to resolve a UDP-address.", Resolve))
		p(fmt.Sprintf(
			"    %d = Failed to listen on a UDP-address.", Listen))
		p(fmt.Sprintf(
			"    %d = Failed to dial a TCP-address.", Dial))
		p("")

		os.Exit(2)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 3 {
		flag.Usage()
	}

	portArg, err := strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
	}
	port = portArg

	publishHost = args[1]

	publishPortArg, err := strconv.Atoi(args[2])
	if err != nil {
		flag.Usage()
	}
	publishPort = publishPortArg
}
