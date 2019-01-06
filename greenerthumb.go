package greenerthumb

import (
	"fmt"
	"os"
	"sort"
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

// KeyError is returned when an expected key is missing from an object.
type KeyError struct {
	Object     map[string]interface{}
	MissingKey string
}

func (e KeyError) Error() string {
	return fmt.Sprintf(`key "%s" is missing from object %v`,
		e.MissingKey, mapToString(e.Object))
}

// StringError is returned when a string is longer than expected.
type StringError struct {
	String string
	Limit  int
}

func (e StringError) Error() string {
	return fmt.Sprintf(
		`string "%s" with length %d is longer than limit %d`,
		e.String, len(e.String), e.Limit)
}

// TypeError is returned when an value has an unexpected type.
type TypeError struct {
	Value interface{}
	Type  string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("value %v has type %T instead of %s",
		e.Value, e.Value, e.Type)
}

func mapToString(x map[string]interface{}) string {
	var keys []string
	for k := range x {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var pairs []string
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s:%v", k, x[k]))
	}
	return fmt.Sprintf("map%v", pairs)
}
