package message

// Disclosure Message.
type Disclosure struct {
	Host string
}

// ID of the Message.
func (m Disclosure) ID() byte {
	return 0x02
}

// Name of the Message.
func (m Disclosure) Name() string {
	return "Disclosure"
}

// SerializeJSON from the Message.
func (m Disclosure) SerializeJSON() map[string]interface{} {
	return map[string]interface{}{"Host": m.Host}
}

// DeserializeJSON into the Message.
func (m *Disclosure) DeserializeJSON(x map[string]interface{}) error {
	if len(x) != 1 {
		return ErrJSON
	}
	xHost, ok := x["Host"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Host",
			Reason: "missing"}
	}
	host, ok := xHost.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "Host",
			Reason: "not a string"}
	}
	m.Host = host
	return nil
}

// SerializeBytes from the Message.
func (m Disclosure) SerializeBytes() []byte {
	bs := []byte{byte(len(m.Host))}
	bs = append(bs, []byte(m.Host)...)
	return bs
}

// DeserializeBytes into the Message.
func (m *Disclosure) DeserializeBytes(bs []byte) error {
	if len(bs) < 1 {
		return ErrBytes
	}
	length := bs[0]
	if byte(len(bs)) < 1+length {
		return ErrBytes
	}
	host := string(bs[1 : 1+length])
	bs = bs[1+length:]

	if len(bs) != 0 {
		return ErrBytes
	}

	m.Host = host

	return nil
}

func init() { addMessage(&Disclosure{}) }
