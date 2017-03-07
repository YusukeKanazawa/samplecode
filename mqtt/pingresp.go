package mqtt

type pingrespMsg struct {
	*fixedHeader
}

func newPingrespMsg() *pingrespMsg {
	return &pingrespMsg{
		fixedHeader: newFixedHeader(PINGRESP),
	}
}

// message interface
func (msg *pingrespMsg) decode(bug []byte) {}
func (msg *pingrespMsg) encode() []byte {
	buff := make([]byte, 0, 2)
	return append(buff, msg.fixedHeader.encode(0)...)
}

func (msg *pingrespMsg) getType() msgType {
	return PINGRESP
}
func (msg *pingrespMsg) String() string {
	return "PINGRESP"
}
