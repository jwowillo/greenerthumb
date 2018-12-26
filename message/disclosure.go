package message

// Disclosure Message.
//
// Name, PublishHost, and CommandHost can't be more than a byte in length.
type Disclosure struct {
	PackageName string
	PublishHost string
	CommandHost string
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
	return map[string]interface{}{
		"PackageName": m.PackageName,
		"PublishHost": m.PublishHost,
		"CommandHost": m.CommandHost}
}

// DeserializeJSON into the Message.
func (m *Disclosure) DeserializeJSON(x map[string]interface{}) error {
	if len(x) != 3 {
		return ErrJSON
	}
	xPackageName, ok := x["PackageName"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "PackageName",
			Reason: "missing"}
	}
	packageName, ok := xPackageName.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "PackageName",
			Reason: "not a string"}
	}
	xPublishHost, ok := x["PublishHost"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "PublishHost",
			Reason: "missing"}
	}
	publishHost, ok := xPublishHost.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "PublishHost",
			Reason: "not a string"}
	}
	xCommandHost, ok := x["CommandHost"]
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "CommandHost",
			Reason: "missing"}
	}
	commandHost, ok := xCommandHost.(string)
	if !ok {
		return JSONError{
			Data:   x,
			BadKey: "CommandHost",
			Reason: "not a string"}
	}
	m.PackageName = packageName
	m.PublishHost = publishHost
	m.CommandHost = commandHost
	return nil
}

// SerializeBytes from the Message.
func (m Disclosure) SerializeBytes() []byte {
	bs := []byte{byte(len(m.PackageName))}
	bs = append(bs, []byte(m.PackageName)...)
	bs = append(bs, byte(len(m.PublishHost)))
	bs = append(bs, []byte(m.PublishHost)...)
	bs = append(bs, byte(len(m.CommandHost)))
	bs = append(bs, []byte(m.CommandHost)...)
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
	packageName := string(bs[1 : 1+length])
	bs = bs[1+length:]

	if len(bs) < 1 {
		return ErrBytes
	}
	length = bs[0]
	if byte(len(bs)) < 1+length {
		return ErrBytes
	}
	publishHost := string(bs[1 : 1+length])
	bs = bs[1+length:]

	if len(bs) < 1 {
		return ErrBytes
	}
	length = bs[0]
	if byte(len(bs)) < 1+length {
		return ErrBytes
	}
	commandHost := string(bs[1 : 1+length])
	bs = bs[1+length:]

	if len(bs) != 0 {
		return ErrBytes
	}

	m.PackageName = packageName
	m.PublishHost = publishHost
	m.CommandHost = commandHost

	return nil
}

func init() { addMessage(&Disclosure{}) }
