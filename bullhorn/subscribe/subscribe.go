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

	"github.com/jwowillo/greenerthumb"
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

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("subscribe", l, args...)
}

func logError(err error) {
	greenerthumb.Error("subscribe", err)
}

func makeConn(writePort int, host string, port int) net.Conn {
	conn, err := net.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil
	}

	realHost, realPort, err := parseHostAndPort(conn.RemoteAddr())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	logInfo("connection to %s:%d started", realHost, realPort)
	fmt.Fprintf(conn, fmt.Sprintf("%d\n", writePort))

	return conn
}

func keepOpen(
	monitorConn, conn net.Conn,
	shouldReconnect bool, delay int,
	port int, publishHost string, publishPort int) error {
	for {
		for monitorConn == nil {
			logInfo(
				"connection to %s:%d unsuccessful",
				publishHost, publishPort)
			if shouldReconnect {
				time.Sleep(time.Duration(delay) * time.Second)

				logInfo(
					"attemping reconnect to %s:%d",
					publishHost, publishPort)

				monitorConn = makeConn(
					port,
					publishHost, publishPort)
			} else {
				return fmt.Errorf(
					"connection to %s:%d failed",
					publishHost, publishPort)
			}
		}

		ioutil.ReadAll(monitorConn)
		monitorConn = nil
	}

	conn.Close()

	return nil
}

func parseHostAndPort(addr net.Addr) (string, int, error) {
	parts := strings.Split(addr.String(), ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return "", -1, err
	}
	return parts[0], port, nil
}

func main() {
	shouldReconnect := reconnectDelay >= 0
	publishHost := host
	publishPort := port

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":0"))
	if err != nil {
		logError(err)
		os.Exit(Resolve)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logError(err)
		os.Exit(Listen)
	}

	_, port, err := parseHostAndPort(conn.LocalAddr())
	if err != nil {
		logError(err)
		os.Exit(ParsePort)
	}

	tcpConn := makeConn(port, publishHost, publishPort)
	go func() {
		if err := keepOpen(
			tcpConn, conn,
			shouldReconnect, reconnectDelay,
			port, publishHost, publishPort); err != nil {

			logError(err)
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
		p("./subscribe <publish_host> <publish_port> \\")
		p("    ?--reconnect-delay <delay>")
		p("")
		p("subscribe to a publisher on a network and write its data to")
		p("STDOUT.")
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
