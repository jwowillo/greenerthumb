package message

// AirStatus Message.
type AirStatus struct {
	Temperature float32
}

// ID of the Message.
func (m AirStatus) ID() byte {
	return 0x00
}

// Name of the Message.
func (m AirStatus) Name() string {
	return "Air"
}

// SerializeJSON from the Message.
func (m AirStatus) SerializeJSON() map[string]interface{} {
	return map[string]interface{}{"Temperature": m.Temperature}
}

// DeserializeJSON into the Message.
func (m *AirStatus) DeserializeJSON(x map[string]interface{}) error {
	if len(x) != 1 {
		return ErrJSON
	}
	xTemperature, ok := x["Temperature"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Temperature",
			Reason: "missing"}
	}
	temperature, ok := xTemperature.(float64)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Temperature",
			Reason: "not a number"}
	}
	m.Temperature = float32(temperature)
	return nil
}

// SerializeBytes from the Message.
func (m AirStatus) SerializeBytes() []byte {
	return floatToBytes(m.Temperature)
}

// DeserializeBytes into the Message.
func (m *AirStatus) DeserializeBytes(bs []byte) error {
	if len(bs) != 4 {
		return ErrBytes
	}
	m.Temperature = bytesToFloat(bs[0:4])
	return nil
}

func init() { addMessage(&AirStatus{}) }
