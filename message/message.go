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
	// ErrBytes is return if bytes Message doesn't have the correct buffer
	// length.
	ErrBytes = errors.New("message doesn't have the correct buffer length")
)

// JSONMessage is a Message that supports the JSON format.
type JSONMessage interface {
	SerializeJSON() map[string]interface{}
	DeserializeJSON(map[string]interface{}) error
}

// BytesMessage is a Message that supports the bytes format.
type BytesMessage interface {
	SerializeBytes() []byte
	DeserializeBytes([]byte) error
}

// Message is an IDed Message that can be converted between bytes and JSON.
type Message interface {
	ID() byte
	Name() string
	JSONMessage
	BytesMessage
}

// Header of a Wrapper.
type Header struct {
	ID        byte
	Name      string
	Timestamp time.Time
	Sender    string
}

// Wrapper of a Message with metadata.
type Wrapper struct {
	Header  Header
	Message Message
}

// ID of the Wrapper.
func (w *Wrapper) ID() byte {
	return w.Header.ID
}

// Name of the Wrapper.
func (w *Wrapper) Name() string {
	return w.Header.Name
}

// SerializeBytes from the Message.
func (w *Wrapper) SerializeBytes() []byte {
	bs := w.Message.SerializeBytes()
	wrapper := make([]byte, 0, 1+8+1+len(w.Header.Sender)+len(bs))
	wrapper = append(wrapper, w.Header.ID)
	wrapper = append(wrapper, intToBytes(w.Header.Timestamp.Unix())...)
	wrapper = append(wrapper, byte(len(w.Header.Sender)))
	wrapper = append(wrapper, []byte(w.Header.Sender)...)
	return append(wrapper, bs...)
}

// DeserializeBytes into the Message.
func (w *Wrapper) DeserializeBytes(bs []byte) error {
	// ID
	if len(bs) < 1 {
		return ErrBytes
	}
	id := bs[0]
	m, ok := messageForID[id]
	if !ok {
		return ErrMessage
	}
	bs = bs[1:]

	// Timestamp
	if len(bs) < 8 {
		return ErrBytes
	}
	t := intToTime(int64(bytesToInt(bs[0:8])))
	bs = bs[8:]

	// Sender Length
	if len(bs) < 1 {
		return ErrBytes
	}
	senderLength := bs[0]
	bs = bs[1:]

	// Sender
	if byte(len(bs)) < senderLength {
		return ErrBytes
	}
	sender := string(bs[:senderLength])
	bs = bs[senderLength:]

	if err := m.DeserializeBytes(bs); err != nil {
		return err
	}

	w.Header.ID = id
	w.Header.Name = m.Name()
	w.Header.Timestamp = t
	w.Header.Sender = sender
	w.Message = m

	return nil
}

// SerializeJSON from the Message.
func (w Wrapper) SerializeJSON() map[string]interface{} {
	s := w.Message.SerializeJSON()
	s["Header"] = map[string]interface{}{
		"Name":      w.Header.Name,
		"Timestamp": timeToInt(w.Header.Timestamp),
		"Sender":    w.Header.Sender,
	}
	return s
}

// DeserializeJSON into the Message.
func (w *Wrapper) DeserializeJSON(x map[string]interface{}) error {
	xHeader, ok := x["Header"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header",
			Reason: "missing"}
	}
	Header, ok := xHeader.(map[string]interface{})
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header",
			Reason: "not a map"}
	}

	xName, ok := Header["Name"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Name",
			Reason: "missing"}
	}
	xTimestamp, ok := Header["Timestamp"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Timestamp",
			Reason: "missing"}
	}
	xSender, ok := Header["Sender"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Sender",
			Reason: "missing"}
	}

	name, ok := xName.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Name",
			Reason: "not a string"}
	}
	timestamp, ok := xTimestamp.(float64)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Timestamp",
			Reason: "not a number"}
	}
	if math.Floor(timestamp) != timestamp {
		return JSONError{
			Data:   x,
			BadKey: "Header/Timestamp",
			Reason: "not an integer"}
	}
	sender, ok := xSender.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Header/Sender",
			Reason: "not a string"}
	}

	delete(x, "Header")

	m, ok := messageForName[name]
	if !ok {
		return ErrMessage
	}
	if err := m.DeserializeJSON(x); err != nil {
		return err
	}

	w.Header.ID = m.ID()
	w.Header.Name = name
	w.Header.Timestamp = intToTime(int64(timestamp))
	w.Header.Sender = sender
	w.Message = m

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
