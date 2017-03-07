package mqtt

type pingreqMsg struct {
	*fixedHeader
}

func newPingreqMsg() *pingreqMsg {
	return &pingreqMsg{
		fixedHeader: newFixedHeader(PINGREQ << 4),
	}
}

// message interface
func (msg *pingreqMsg) decode(buffer []byte) {

}
func (msg *pingreqMsg) encode() []byte {
	return nil //&[PINGREQ << 4, 0x00]
}
func (msg *pingreqMsg) getType() msgType {
	return PINGREQ
}
func (msg *pingreqMsg) String() string {
	return "PINGREQ"
}
