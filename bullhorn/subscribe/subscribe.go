package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_ = iota
	_
	// Resolve is the error-code for failing to resolve a UDP-address.
	Resolve = 1 << iota
	// Listen is the error-code for failing to listen on a UDP-address.
	Listen = 1 << iota
	// Connect is the error-code for failing to connect to a TCP-address.
	Connect = 1 << iota
	// ParsePort is the error-code for failing to parse a port from a
	// connection.
	ParsePort = 1 << iota
)

func makeConn(writePort int, host string, port int) net.Conn {
	conn, err := net.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil
	}
	fmt.Fprintln(os.Stderr, "connection successful")
	fmt.Fprintf(conn, fmt.Sprintf("%d\n", writePort))
	return conn
}

func keepOpen(
	monitorConn, conn net.Conn,
	shouldReconnect bool, delay int,
	port int, publishHost string, publishPort int) error {
	for {
		for monitorConn == nil {
			if shouldReconnect {
				time.Sleep(time.Duration(delay) * time.Second)

				fmt.Fprintln(os.Stderr, "attempting reconnect")

				monitorConn = makeConn(
					port,
					publishHost, publishPort)
			} else {
				return fmt.Errorf(
					"connect to %s:%d failed",
					publishHost, publishPort)
			}
		}

		ioutil.ReadAll(monitorConn)
		monitorConn = nil
	}

	conn.Close()

	return nil
}

func parsePort(conn net.Conn) (int, error) {
	parts := strings.Split(conn.LocalAddr().String(), ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return -1, err
	}
	return port, nil
}

func main() {
	shouldReconnect := reconnectDelay >= 0
	publishHost := host
	publishPort := port

	handler := func(err error) { fmt.Fprintln(os.Stderr, err) }
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":0"))
	if err != nil {
		handler(err)
		os.Exit(Resolve)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		handler(err)
		os.Exit(Listen)
	}

	port, err := parsePort(conn)
	if err != nil {
		handler(err)
		os.Exit(ParsePort)
	}

	tcpConn := makeConn(port, publishHost, publishPort)
	go func() {
		if err := keepOpen(
			tcpConn, conn,
			shouldReconnect, reconnectDelay,
			port, publishHost, publishPort); err != nil {
			handler(err)
			os.Exit(Connect)
		}
	}()

	io.Copy(os.Stdout, conn)
}

// host of the publisher.
var host string

// port of the publisher.
var port int

// reconnectDelay is the delay before attemping a reconnect when the connection
// is lost.
var reconnectDelay int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("subscribe to a publisher on a network and write its data to")
		p("STDOUT as a subscriber.")
		p("")
		p("The publisher's host and port must be passed. A reconnect")
		p("delay that will cause the subscriber to attempt to")
		p("reconnect to the publisher can also be passed.")
		p("")
		p("An example that connects to a publisher and attempts to")
		p("reconnect every 5 seconds when a connection is lost is:")
		p("")
		p("    ./subscribe 127.0.0.1 5050 --reconnect-delay 5")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to resolve a UDP-address.", Resolve))
		p(fmt.Sprintf(
			"    %d = Failed to listen on a UDP-address.", Listen))
		p(fmt.Sprintf(
			"    %d = Failed to connect to a TCP-address.",
			Connect))
		p(fmt.Sprintf(
			"    %d = Failed to parse a port from a connection.",
			ParsePort))
		p("")

		os.Exit(2)
	}

	flag.IntVar(&reconnectDelay, "reconnect-delay", -1,
		"delay in seconds before attempting a reconnect")

	if len(os.Args) < 3 {
		flag.Usage()
	}

	flag.CommandLine.Parse(os.Args[3:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}

	host = os.Args[1]

	portArg, err := strconv.Atoi(os.Args[2])
	if err != nil {
		flag.Usage()
	}
	port = portArg
}
