package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/jwowillo/greenerthumb"
)

const (
	_ = iota
	_
	// Listen is the error-code for failing to listen for TCP-connections.
	Listen = 1 << iota
	// Accept is the error-code for failing to accept a TCP-connection.
	Accept = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("bullhorn-publish", l, args...)
}

func logError(err error) {
	greenerthumb.Error("bullhorn-publish", err)
}

// conns are active net.Conns.
var mux sync.RWMutex
var conns map[string]net.Conn = make(map[string]net.Conn)

func parseHost(addr net.Addr) string {
	parts := strings.Split(addr.String(), ":")
	return strings.Join(parts[:len(parts)-1], "")
}

func readPort(conn net.Conn) (int, error) {
	portString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(portString[:len(portString)-1])
}

// storeUntilClosed stores the UDP connection indicated from the net.Conn until
// the net.Conn is closed.
func storeUntilClosed(conn net.Conn) error {
	defer conn.Close()
	host := parseHost(conn.RemoteAddr())
	port, err := readPort(conn)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	udpConn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	func() {
		logInfo("connection to %s:%d started", host, port)
		mux.Lock()
		defer mux.Unlock()
		conns[addr] = udpConn
	}()
	ioutil.ReadAll(conn)
	func() {
		defer conns[addr].Close()
		logInfo("connection to %s:%d ended", host, port)
		mux.Lock()
		defer mux.Unlock()
		delete(conns, addr)
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
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logError(err)
		os.Exit(Listen)
	}
	defer ln.Close()
	logInfo("listening at :%d", port)

	go acceptConnections(ln, logError)
	if err := readIntoConns(os.Stdin, logError); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

// port to listen for TCP-connections on.
var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./publish <port>")
		p("")
		p("publish messages from STDIN on a network to subscribers.")
		p("")
		p("A port to listen on for subscribers must be passed.")
		p("")
		p("An example that publishes 'a' and 'b' to subscribers is:")
		p("")
		p("    ./publish 5050")
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
			"    %d = Failed to accept a TCP-connection.",
			Accept))
		p(fmt.Sprintf(
			"    %d = Failed to read input.",
			ReadInput))
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
