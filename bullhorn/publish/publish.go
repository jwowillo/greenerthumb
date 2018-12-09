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
)

// mux synchronizes net.Conns.
var mux sync.RWMutex

// conns are active net.Conns.
var conns map[string]net.Conn = make(map[string]net.Conn)

// Host part of the address.
func Host(addr net.Addr) string {
	parts := strings.Split(addr.String(), ":")
	return strings.Join(parts[:len(parts)-1], "")
}

// ReadPort from the net.Conn.
func ReadPort(conn net.Conn) (int, error) {
	portString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(portString[:len(portString)-1])
}

// StoreUntilClosed stores the UDP connection indicated from the net.Conn until
// the net.Conn is closed.
func StoreUntilClosed(conn net.Conn) error {
	defer conn.Close()
	host := Host(conn.RemoteAddr())
	port, err := ReadPort(conn)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	udpConn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	func() {
		mux.Lock()
		defer mux.Unlock()
		conns[addr] = udpConn
	}()
	ioutil.ReadAll(conn)
	func() {
		mux.Lock()
		defer mux.Unlock()
		delete(conns, addr)
	}()
	return nil
}

// ReadIntoConns reads lines from the io.Reader into all the stored net.Conns.
func ReadIntoConns(r io.Reader, handler func(error)) {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			handler(err)
			continue
		}
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
}

// AcceptConnections accepts connections from the net.Listener until closed.
func AcceptConnections(ln net.Listener, handler func(err error)) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			handler(err)
			return
		}
		go func() {
			if err := StoreUntilClosed(conn); err != nil {
				handler(err)
			}
		}()
	}
}

const (
	_ = iota
	_
	// Listen is the error-code for failing to listen for TCP-connections.
	Listen = 1 << iota
	// Accept is the error-code for failing to accept a TCP-connection.
	Accept = 1 << iota
)

func main() {
	handler := func(err error) { fmt.Fprintln(os.Stderr, err) }
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		handler(err)
		os.Exit(Listen)
	}
	go AcceptConnections(ln, handler)
	ReadIntoConns(os.Stdin, handler)
}

// port to listen for TCP-connections on.
var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("publish data from STDIN on a network to bullhorn")
		p("subscribers as a bullhorn publisher.")
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
