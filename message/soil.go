package message

// SoilStatus Message.
type SoilStatus struct {
	Moisture float32
}

// ID of the Message.
func (m SoilStatus) ID() byte {
	return 0x01
}

// Name of the Message.
func (m SoilStatus) Name() string {
	return "Soil"
}

// SerializeJSON from the Message.
func (m SoilStatus) SerializeJSON() map[string]interface{} {
	return map[string]interface{}{"Moisture": m.Moisture}
}

// DeserializeJSON into the Message.
func (m *SoilStatus) DeserializeJSON(x map[string]interface{}) error {
	if len(x) != 1 {
		return ErrJSON
	}
	xMoisture, ok := x["Moisture"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Moisture",
			Reason: "missing"}
	}
	moisture, ok := xMoisture.(float64)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Moisture",
			Reason: "not a number"}
	}
	m.Moisture = float32(moisture)
	return nil
}

// SerializeBytes from the Message.
func (m SoilStatus) SerializeBytes() []byte {
	return floatToBytes(m.Moisture)
}

// DeserializeBytes into the Message.
func (m *SoilStatus) DeserializeBytes(bs []byte) error {
	if len(bs) != 4 {
		return ErrBytes
	}
	m.Moisture = bytesToFloat(bs)
	return nil
}

func init() { addMessage(&SoilStatus{}) }
