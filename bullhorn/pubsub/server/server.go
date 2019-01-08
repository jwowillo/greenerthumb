package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sync"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// Listen is the error-code for failing to listen for TCP-connections.
	Listen = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("bullhorn-pubsub-server", l, args...)
}

func logError(err error) {
	greenerthumb.Error("bullhorn-pubsub-server", err)
}

// conns are active net.Conns.
var mux sync.RWMutex
var conns map[string]net.Conn = make(map[string]net.Conn)

func readHost(conn net.Conn) (string, error) {
	host, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return host[:len(host)-1], nil
}

// storeUntilClosed stores the UDP connection indicated from the net.Conn until
// the net.Conn is closed.
func storeUntilClosed(conn net.Conn) error {
	defer conn.Close()
	host, err := readHost(conn)
	if err != nil {
		return err
	}

	udpConn, err := net.Dial("udp", host)
	if err != nil {
		return err
	}
	func() {
		logInfo("connection to %s started", host)
		mux.Lock()
		defer mux.Unlock()
		conns[host] = udpConn
	}()
	ioutil.ReadAll(conn)
	func() {
		defer conns[host].Close()
		logInfo("connection to %s ended", host)
		mux.Lock()
		defer mux.Unlock()
		delete(conns, host)
	}()
	return nil
}

// readIntoConns reads lines from the io.Reader into all the stored net.Conns.
func readIntoConns(r io.Reader, handler func(error)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := append(scanner.Bytes(), '\n')

		func() {
			mux.RLock()
			defer mux.RUnlock()
			for _, conn := range conns {
				if _, err := conn.Write(line); err != nil {
					handler(err)
				}
			}
		}()
	}

	return scanner.Err()
}

// acceptConnections accepts connections from the net.Listener until closed.
func acceptConnections(ln net.Listener, handler func(err error)) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			handler(err)
			return
		}
		go func() {
			if err := storeUntilClosed(conn); err != nil {
				handler(err)
			}
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

	go acceptConnections(ln, logError)
	if err := readIntoConns(os.Stdin, logError); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

var host string

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./server ?--host <host>")
		p("")
		p("server publishes messages from STDIN on a network to")
		p("subscribers.")
		p("")
		p("The host to listen for connections on can be optionally")
		p("passed. The default is ':0' which chooses a random port.")
		p("")
		p("An example that publishes 'a' and 'b' to subscribers is:")
		p("")
		p("    ./server")
		p("")
		p("    < a")
		p("    < b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to listen for TCP-connections.",
			Listen))
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
		p("")

		os.Exit(2)
	}

	flag.StringVar(&host, "host", ":0", "host to publish at")
	flag.Parse()

	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
