package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// Listen is the error-code for failing to listen for TCP-connections.
	Listen = 1 << iota
	// Accept is the error-code for failing to accept a TCP-connection.
	Accept = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("bullhorn-listen", l, args...)
}

func logError(err error) {
	greenerthumb.Error("bullhorn-listen", err)
}

func acceptConnections(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		logInfo("connection to %s started", conn.RemoteAddr())
		go func() {
			defer conn.Close()
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				logError(err)
			}
			logInfo("connection to %s ended", conn.RemoteAddr())
		}()
	}
}

func main() {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		logError(err)
		os.Exit(Listen)
	}
	defer ln.Close()
	logInfo("listening at %s", ln.Addr())

	if err := acceptConnections(ln); err != nil {
		logError(err)
		os.Exit(Accept)
	}
}

var host string

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./listen ?--host <host>")
		p("")
		p("listen to talkers by receiving their TCP messages and")
		p("writing them to STDOUT.")
		p("")
		p("The host to listen for connections on can be optionally")
		p("passed. The default is ':0' which chooses a random port.")
		p("")
		p("An example that receives 'a' and 'b' from talkers is:")
		p("")
		p("    ./listen")
		p("")
		p("    a")
		p("    b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to listen for TCP-connections.",
			Listen))
		p(fmt.Sprintf(
			"    %d = Failed to accept a TCP-connection.",
			Accept))
		p("")

		os.Exit(2)
	}
	flag.StringVar(&host, "host", ":0", "host to listen at")
	flag.Parse()

	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
