package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jwowillo/greenerthumb"
)

// Header of a message.
type Header map[string]interface{}

// GetInt from the Header.
//
// Return an error if the value is bad.
func (h Header) GetInt(k string) (int64, error) {
	x, ok := h[k]
	if !ok {
		return -1, greenerthumb.KeyError{
			Object:     h,
			MissingKey: "Header/" + k}
	}
	v, ok := x.(float64)
	if !ok {
		return -1, greenerthumb.TypeError{Value: x, Type: "float64"}
	}
	return int64(v), nil
}

// GetString from the Header.
//
// Return an error if the value is bad.
func (h Header) GetString(k string) (string, error) {
	x, ok := h[k]
	if !ok {
		return "", greenerthumb.KeyError{
			Object:     h,
			MissingKey: "Header/" + k}
	}
	v, ok := x.(string)
	if !ok {
		return "", greenerthumb.TypeError{Value: x, Type: "string"}
	}
	return v, nil
}

// Serialize the Header, field, and value into a JSON string.
func Serialize(h Header, field string, value float64) (string, error) {
	bs, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`{"Header":%s,"%s":%f}`, bs, field, value), nil
}

// FieldHandler is called with the Header for the field and the field
// name and value.
//
// Returns an error if the field couldn't be processed.
type FieldHandler func(Header, string, float64)

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
		xHeader, ok := x["Header"]
		if !ok {
			return greenerthumb.KeyError{
				Object:     x,
				MissingKey: "Header"}
		}
		header, ok := xHeader.(map[string]interface{})
		if !ok {
			return greenerthumb.TypeError{
				Value: xHeader,
				Type:  "Header"}
		}

		delete(x, "Header")

		for k, rawV := range x {
			v, ok := rawV.(float64)
			if !ok {
				return greenerthumb.TypeError{
					Value: rawV,
					Type:  "float64"}
			}
			cb(header, k, v)
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
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		if err := cb(scanner.Bytes()); err != nil {
			ecb(err)
		}
	}
	return scanner.Err()
}
