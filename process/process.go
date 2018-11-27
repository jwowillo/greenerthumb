package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// FieldHandler is called with the message name and timestamp for the field and
// the field name and value.
//
// Returns an error if the field couldn't be processed.
type FieldHandler func(string, int64, string, float64)

//ErrorHandler is called whenever an error occurs.
type ErrorHandler func(error)

// Fields parses fields from each JSON line in the io.Reader and passes them to
// the handler.
//
// Passes errors with data to the ErrorHandler.
//
// Returns an error if the io.Reader can't be read.
func Fields(rd io.Reader, cb FieldHandler, ecb ErrorHandler) error {
	f := func(x map[string]interface{}) error {
		rawName, ok := x["Name"]
		if !ok {
			return KeyError{Object: x, MissingKey: "Name"}
		}
		rawTimestamp, ok := x["Timestamp"]
		if !ok {
			return KeyError{Object: x, MissingKey: "Timestamp"}
		}
		delete(x, "Name")
		delete(x, "Timestamp")

		name, ok := rawName.(string)
		if !ok {
			return TypeError{Value: rawName, Type: "string"}
		}
		timestamp, ok := rawTimestamp.(float64)
		if !ok {
			return TypeError{Value: rawTimestamp, Type: "float64"}
		}

		for k, rawV := range x {
			v, ok := rawV.(float64)
			if !ok {
				return TypeError{Value: rawV, Type: "float64"}
			}
			cb(name, int64(timestamp), k, v)
		}
		return nil
	}

	return parseObjects(rd, f, ecb)
}

func parseObjects(
	rd io.Reader,
	cb func(map[string]interface{}) error, ecb ErrorHandler) error {
	f := func(bs []byte) error {
		var x map[string]interface{}
		if err := json.Unmarshal(bs, &x); err != nil {
			return err
		}
		return cb(x)
	}
	return parseLines(rd, f, ecb)
}

func parseLines(rd io.Reader, cb func([]byte) error, ecb ErrorHandler) error {
	scanner := bufio.NewScanner(bufio.NewReader(rd))
	for scanner.Scan() {
		if err := cb(scanner.Bytes()); err != nil {
			ecb(err)
		}
	}
	return scanner.Err()
}

// KeyError is returned when an expected key is missing from an object.
type KeyError struct {
	Object     map[string]interface{}
	MissingKey string
}

func (e KeyError) Error() string {
	return fmt.Sprintf("key \"%s\" is missing from object %v",
		e.MissingKey, e.Object)
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
