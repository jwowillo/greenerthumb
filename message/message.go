package message

import (
	"errors"
	"math"
	"time"
)

var (
	// ErrMessage is returned if a Message doesn't exist.
	ErrMessage = errors.New("message doesn't exist")
	// ErrJSON is returned if a Message doesn't have the right amount of
	// keys.
	ErrJSON = errors.New("message doesn't have the correct number of keys")
	// ErrBytes is return if bytes message doesn't have the correct buffer
	// length.
	ErrBytes = errors.New("message doesn't have the correct buffer length")
)

// JSONMessage is a message that supports the JSON format.
type JSONMessage interface {
	SerializeJSON() map[string]interface{}
	DeserializeJSON(map[string]interface{}) error
}

// BytesMessage is a message that supports the bytes format.
type BytesMessage interface {
	SerializeBytes() []byte
	DeserializeBytes([]byte) error
}

// Message is an IDed message that can be converted between bytes and JSON.
type Message interface {
	ID() byte
	Name() string
	JSONMessage
	BytesMessage
}

// Wrapper of a Message with metadata.
type Wrapper struct {
	id        byte
	name      string
	timestamp time.Time
	message   Message
}

// ID of the Wrapper.
func (w *Wrapper) ID() byte {
	return w.id
}

// Name of the Wrapper.
func (w *Wrapper) Name() string {
	return w.name
}

// Timestamp of the Wrapper.
func (w *Wrapper) Timestamp() time.Time {
	return w.timestamp
}

// Message of the Wrapper.
func (w *Wrapper) Message() Message {
	return w.message
}

// SerializeBytes from the Message.
func (w *Wrapper) SerializeBytes() []byte {
	bs := w.message.SerializeBytes()
	wrapper := make([]byte, 0, 1+8+len(bs)+1)
	wrapper = append(wrapper, w.id)
	wrapper = append(wrapper, intToBytes(w.timestamp.Unix())...)
	wrapper = append(wrapper, bs...)
	var sum byte
	for _, b := range wrapper {
		sum += b
	}
	return append(wrapper, sum)
}

// DeserializeBytes into the Message.
func (w *Wrapper) DeserializeBytes(bs []byte) error {
	// ID + Timestamp + Checksum
	if len(bs) < 1+8+1 {
		return ErrBytes
	}
	actualSum := bs[len(bs)-1]
	var expectedSum byte
	for i := 0; i < len(bs)-1; i++ {
		expectedSum += bs[i]
	}
	if actualSum != expectedSum {
		return ErrBytes
	}
	t := intToTime(int64(bytesToInt(bs[1:9])))
	id := bs[0]
	m, ok := messageForID[id]
	if !ok {
		return ErrMessage
	}
	if err := m.DeserializeBytes(bs[9 : len(bs)-1]); err != nil {
		return err
	}
	w.id = id
	w.name = m.Name()
	w.timestamp = t
	w.message = m
	return nil
}

// SerializeJSON from the Message.
func (w Wrapper) SerializeJSON() map[string]interface{} {
	s := w.message.SerializeJSON()
	s["Name"] = w.name
	s["Timestamp"] = timeToInt(w.timestamp)
	return s
}

// DeserializeJSON into the Message.
func (w *Wrapper) DeserializeJSON(x map[string]interface{}) error {
	xName, ok := x["Name"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Name",
			Reason: "missing"}
	}
	xTimestamp, ok := x["Timestamp"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Timestamp",
			Reason: "missing"}
	}
	name, ok := xName.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Name",
			Reason: "not a string"}
	}
	timestamp, ok := xTimestamp.(float64)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Timestamp",
			Reason: "not a number"}
	}
	if math.Floor(timestamp) != timestamp {
		return JSONError{
			Data:   x,
			BadKey: "Timestamp",
			Reason: "not an integer"}
	}
	delete(x, "Name")
	delete(x, "Timestamp")
	m, ok := messageForName[name]
	if !ok {
		return ErrMessage
	}
	if err := m.DeserializeJSON(x); err != nil {
		return err
	}
	w.id = m.ID()
	w.name = name
	w.timestamp = intToTime(int64(timestamp))
	w.message = m
	return nil
}

var (
	messageForID   = map[byte]Message{}
	messageForName = map[string]Message{}
)

func addMessage(m Message) {
	messageForID[m.ID()] = m
	messageForName[m.Name()] = m
}
