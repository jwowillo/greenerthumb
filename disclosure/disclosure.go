package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func serialize(deviceName, publishHost, commandHost string) string {
	return fmt.Sprintf(
		`{"Name":"%s","Timestamp":%d,"DeviceName":"%s","PublishHost":"%s","CommandHost":"%s"}`,
		"Disclosure", time.Now().Unix(),
		deviceName, publishHost, commandHost)
}

func main() {
	duration := time.Duration(float64(time.Second) / rate)
	for {
		fmt.Println(serialize(deviceName, publishHost, commandHost))
		time.Sleep(duration)
	}
}

var (
	deviceName  string
	publishHost string
	commandHost string
)

var (
	rate float64
)

func init() {
	p := func(l string) { fmt.Fprintln(os.Stderr, l) }
	flag.Usage = func() {
		p(`./disclosure <device_name> <publish_host> <command_host> \`)
		p("    ?--rate <rate>")
		p("")
		p("disclosure prints the disclosure message with the passed")
		p("values periodically at the given rate in hertz.")
		p("")
		p("The default rate is 5 hertz.")
		p("")
		p("An example after one second is:")
		p("")
		p("    ./disclosure device :8080 :8081 --rate 1")
		p("")
		p(`    {"Name":"disclosure","Timestamp":0,"DeviceName":"device","PublishHost":":8080","CommandHost":":8081"}`)
		p("")

		os.Exit(2)
	}

	flag.Float64Var(
		&rate,
		"rate",
		5,
		"rate to print disclosures at")

	if len(os.Args) < 4 {
		flag.Usage()
	}

	deviceName = os.Args[1]
	publishHost = os.Args[2]
	commandHost = os.Args[3]

	flag.CommandLine.Parse(os.Args[4:])

	if len(flag.Args()) != 0 {
		flag.Usage()
	}
}
