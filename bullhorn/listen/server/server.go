package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/bullhorn"
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
	greenerthumb.Info("greenerthumb-bullhorn-listen-server", l, args...)
}

func logError(err error) {
	greenerthumb.Error("greenerthumb-bullhorn-listen-server", err)
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

			for {
				bs, err := bullhorn.ReadLength(conn)
				if err != nil {
					if err == io.EOF {
						break
					} else {
						logError(err)
					}
				}

				fmt.Printf("%s\n", greenerthumb.BytesToHex(bs))
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
		p("./server ?--host <host>")
		p("")
		p("server listens to talkers by receiving their TCP messages")
		p("and writing them to STDOUT.")
		p("")
		p("The host to listen for connections on can be optionally")
		p("passed. The default is ':0' which chooses a random port.")
		p("")
		p("An example that receives 'a' and 'b' from clients is:")
		p("")
		p("    ./server")
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
