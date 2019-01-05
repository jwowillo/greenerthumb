package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func serialize(host string) string {
	return fmt.Sprintf(
		`{"Name":"%s","Timestamp":%d,"Host":"%s"}`,
		"Disclosure", time.Now().Unix(),
		host)
}

func main() {
	duration := time.Duration(float64(time.Second) / rate)
	for {
		fmt.Println(serialize(host))
		time.Sleep(duration)
	}
}

var host string

var rate float64

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p(`./disclosure <host> ?--rate <rate>`)
		p("")
		p("disclosure prints the disclosure message with the passed")
		p("values periodically at the given rate in hertz.")
		p("")
		p("The default rate is 5 hertz.")
		p("")
		p("An example after one second is:")
		p("")
		p("    ./disclosure :8080 --rate 1")
		p("")
		p(`    {"Name":"Disclosure","Timestamp":0,"Host":":8080"}`)
		p("")

		os.Exit(2)
	}

	flag.Float64Var(
		&rate,
		"rate",
		5,
		"rate to print disclosures at")

	if len(os.Args) < 2 {
		flag.Usage()
	}

	host = os.Args[1]

	flag.CommandLine.Parse(os.Args[2:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
