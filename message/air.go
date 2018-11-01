package message

// AirStatus Message.
type AirStatus struct {
	Temperature, Humidity float32
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
	return map[string]interface{}{
		"Temperature": m.Temperature,
		"Humidity":    m.Humidity}
}

// DeserializeJSON into the Message.
func (m *AirStatus) DeserializeJSON(x map[string]interface{}) error {
	if len(x) != 2 {
		return ErrJSON
	}
	xTemperature, ok := x["Temperature"]
	if !ok {
		return ErrJSON
	}
	xHumidity, ok := x["Humidity"]
	if !ok {
		return ErrJSON
	}
	temperature, ok := xTemperature.(float64)
	if !ok {
		return ErrJSON
	}
	humidity, ok := xHumidity.(float64)
	if !ok {
		return ErrJSON
	}
	m.Temperature = float32(temperature)
	m.Humidity = float32(humidity)
	return nil
}

// SerializeBytes from the Message.
func (m AirStatus) SerializeBytes() []byte {
	return append(floatToBytes(m.Temperature), floatToBytes(m.Humidity)...)
}

// DeserializeBytes into the Message.
func (m *AirStatus) DeserializeBytes(bs []byte) error {
	if len(bs) != 8 {
		return ErrBytes
	}
	m.Temperature = bytesToFloat(bs[0:4])
	m.Humidity = bytesToFloat(bs[4:8])
	return nil
}

func init() { addMessage(&AirStatus{}) }
