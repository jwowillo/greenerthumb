package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/bullhorn"
)

const (
	_ = iota
	_
	// Dial is the error-code for failing to dial a TCP address.
	Dial = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("greenerthumb-bullhorn-listen-client", l, args...)
}

func logError(err error) {
	greenerthumb.Error("greenerthumb-bullhorn-listen-client", err)
}

func main() {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		logError(err)
		os.Exit(Dial)
	}
	defer conn.Close()
	logInfo("connection to %s started", conn.RemoteAddr())

	done := make(chan interface{})
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			bs, err := greenerthumb.HexToBytes(scanner.Bytes())
			if err != nil {
				logError(err)
				continue
			}

			bs, err = bullhorn.AddLength(bs)
			if err != nil {
				logError(err)
				continue
			}

			conn.Write(bs)
		}

		if err := scanner.Err(); err != nil {
			logError(err)
			os.Exit(ReadInput)
		}

		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()
	go func() {
		ioutil.ReadAll(conn)
		done <- struct{}{}
	}()
	<-done
}

var host string

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./client <host>")
		p("")
		p("client sends messages to servers via TCP.")
		p("")
		p("An example that sends 'A' and 'B' is:")
		p("")
		p("    ./client :5050")
		p("")
		p("    < a")
		p("    < b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to dial a TCP address.",
			Dial))
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

	host = args[0]
}
