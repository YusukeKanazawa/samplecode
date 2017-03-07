package mqtt

type caRC byte

const (
	CONNACK_OK caRC = iota
	CONNACK_PROTOOCL
	CONNACK_ID
	CONNACK_NOT_AVAILABLE
	CONNACK_USER_PASS
	CONNACK_AUTH
)

type connackMsg struct {
	*fixedHeader
	isCleanSession bool
	returnCode caRC
}

func newConnackMsg(rc caRC, cleanSession bool) *connackMsg{
	return &connackMsg{
		fixedHeader: newFixedHeader(CONNACK),
		isCleanSession: cleanSession,
		returnCode: rc,
	}
}

func (m *connackMsg) decode(b []byte) {}

// payload
//              Session PresentFlag
//   1:         |
//      0000 000X
//   2: return code
func (msg *connackMsg) encode() []byte {
	buff := make([]byte, 0, 4)
	buff = append(buff, msg.fixedHeader.encode(2)...)
	if msg.isCleanSession {
		buff = append(buff, 0x01)
	}else{
		buff = append(buff, 0x00)
	}
	return append(buff, byte(msg.returnCode))
}


func (m *connackMsg) getType() msgType{
	return CONNACK
}
