package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
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
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("bullhorn-pubsub-client", l, args...)
}

func logError(err error) {
	greenerthumb.Error("bullhorn-pubsub-client", err)
}

func makeConn(writeHost string, host string) net.Conn {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil
	}

	realHost := conn.RemoteAddr().String()

	logInfo("connection to %s started", realHost)
	fmt.Fprintf(conn, fmt.Sprintf("%s\n", writeHost))

	return conn
}

func keepOpen(
	monitorConn, conn net.Conn,
	shouldReconnect bool, delay int,
	writeHost, publishHost string) error {
	for {
		for monitorConn == nil {
			logInfo("connection to %s unsuccessful", publishHost)
			if shouldReconnect {
				time.Sleep(time.Duration(delay) * time.Second)

				logInfo("attemping reconnect to %s",
					publishHost)

				monitorConn = makeConn(writeHost, publishHost)
			} else {
				return fmt.Errorf(
					"connection to %s failed",
					publishHost)
			}
		}

		ioutil.ReadAll(monitorConn)
		monitorConn.Close()
		monitorConn = nil
	}

	return nil
}

func main() {
	shouldReconnect := reconnectDelay >= 0
	publishHost := host

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
	defer conn.Close()

	writeHost := conn.LocalAddr().String()
	tcpConn := makeConn(writeHost, publishHost)
	if tcpConn != nil {
		defer tcpConn.Close()
	}
	go func() {
		if err := keepOpen(
			tcpConn, conn, shouldReconnect, reconnectDelay,
			writeHost, publishHost); err != nil {

			logError(err)
			os.Exit(Connect)
		}
	}()

	io.Copy(os.Stdout, conn)
}

// host of the publisher.
var host string

// reconnectDelay is the delay before attemping a reconnect when the connection
// is lost.
var reconnectDelay int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./client <publish_host> ?--reconnect-delay <delay>")
		p("")
		p("client subscribes to a publisher on a network and writes")
		p("its messages to STDOUT.")
		p("")
		p("The publisher's host must be passed. A reconnect delay that")
		p("will cause the subscriber to attempt to reconnect to the")
		p("publisher can also be passed.")
		p("")
		p("An example that connects to a publisher and attempts to")
		p("reconnect every 5 seconds when a connection is lost is:")
		p("")
		p("    ./client 127.0.0.1:5050 --reconnect-delay 5")
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
		p("")

		os.Exit(2)
	}

	flag.IntVar(&reconnectDelay, "reconnect-delay", -1,
		"delay in seconds before attempting a reconnect")

	if len(os.Args) < 2 {
		flag.Usage()
	}

	flag.CommandLine.Parse(os.Args[2:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}

	host = os.Args[1]
}
