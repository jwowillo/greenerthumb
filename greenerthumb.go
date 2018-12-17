package greenerthumb

import (
	"fmt"
	"os"
	"time"
)

func log(mode, program, l string, args ...interface{}) {
	f := fmt.Sprintf(
		"%s %s %s - %s\n",
		mode,
		program,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		l)
	fmt.Fprintf(os.Stderr, f, args...)
}

// Info logs an info message with arguments.
func Info(program, l string, args ...interface{}) {
	log("INFO", program, l, args...)
}

// Error logs an error.
func Error(program string, err error) {
	log("ERROR", program, err.Error())
}
