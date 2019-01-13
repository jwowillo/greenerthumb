package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/jwowillo/greenerthumb"
	"github.com/jwowillo/greenerthumb/bullhorn"
)

const (
	_ = iota
	_
	// Dial is the error-code for failing to dial a UDP-connections.
	Dial = 1 << iota
	// ReadInput is the error-code for failing to read input.
	ReadInput = 1 << iota
)

func logInfo(l string, args ...interface{}) {
	greenerthumb.Info("greenerthumb-bullhorn-broadcast-server", l, args...)
}

func logError(err error) {
	greenerthumb.Error("greenerthumb-bullhorn-broadcast-server", err)
}

func main() {
	conn, err := net.Dial("udp", fmt.Sprintf("255.255.255.255:%d", port))
	if err != nil {
		logError(err)
		os.Exit(Dial)
	}
	defer conn.Close()

	logInfo("broadcasting on port %d", port)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bs, err := greenerthumb.HexToBytes(scanner.Bytes())
		if err != nil {
			logError(err)
			continue
		}

		bs = append(bullhorn.Uint32ToBytes(bullhorn.Sum(bs)), bs...)

		conn.Write(bs)
	}

	if err := scanner.Err(); err != nil {
		logError(err)
		os.Exit(ReadInput)
	}
}

var port int

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p("")
		p("./server <port>")
		p("")
		p("server broadcasts messages to all clients.")
		p("")
		p("An example that broadcasts 'a' and 'b ' is:")
		p("")
		p("    ./server 5050")
		p("")
		p("    < a")
		p("    < b")
		p("")
		p("Error-codes are used for the following:")
		p("")
		p(fmt.Sprintf(
			"    %d = Failed to dial a UDP-connection.",
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

	portArg, err := strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
	}
	port = portArg
}
