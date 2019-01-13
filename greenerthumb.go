package greenerthumb

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

// ErrBytes is returned when a byte-string has odd-length.
var ErrBytes = errors.New("byte-string has odd-length")

// HexToBytes converts base-16 to actual bytes.
//
// Returns ErrBytes if the input string has odd-length.
func HexToBytes(bs []byte) ([]byte, error) {
	if len(bs)%2 != 0 {
		return nil, ErrBytes
	}
	converted := make([]byte, len(bs)/2)
	for i := 0; i < len(converted); i++ {
		converted[i] = (toByte(bs[i*2]) << 4) | toByte(bs[i*2+1])
	}

	return converted, nil
}

// BytesToHex converts raw bytes to base-16.
func BytesToHex(bs []byte) []byte {
	buff := &bytes.Buffer{}
	for _, b := range bs {
		fmt.Fprintf(buff, "%02x", b)
	}
	return buff.Bytes()
}

// toByte converts a hex-character to its byte value.
func toByte(x byte) byte {
	switch x {
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		return x - 0x30
	case 'a':
		fallthrough
	case 'b':
		fallthrough
	case 'c':
		fallthrough
	case 'd':
		fallthrough
	case 'e':
		fallthrough
	case 'f':
		return x - (0x66 - 0xf)
	}
	return 0
}

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
