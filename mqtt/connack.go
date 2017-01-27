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
	returnCode caRC
}

func newConnackMsg(rc caRC) *connackMsg{
	return &connackMsg{
		returnCode: rc,
	}
}

func (m *connackMsg) decode(b []byte) {
}

// payload
//   1: reserved
//   2: return code
func (m *connackMsg) encode() []byte {
	buff := make([]byte, 0, 4)
	buff = append(buff, m.fixedHeader.encode(2)...)
	return append(buff, 0x00, byte(m.returnCode))
}


func (m *connackMsg) proc(){
}
