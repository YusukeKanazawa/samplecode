package mqtt

import (
	"log"
)

const (
	RESERVED0 msgType = iota
	CONNECT
	CONNACK
	PUBLISH
	PUBACK
	PUBREC
	PUBREL
	PUBCOMP
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PINGREQ
	PINGRESP
	DISCONNECT
	RESERVED15
)

type msgType uint8

type message interface {
	decode([]byte)
	encode() []byte
}

func NewMessage(b []byte) message {
	var msg message
	head, l := decodeFixedHeader(b)

	switch head.msgType {
	case RESERVED0:
	case CONNECT: // client -> server
		con := newConnectMsg()
    con.fixedHeader = head
    msg = con
		log.Println("->Connect")
	case CONNACK: // server -> client
	case PUBLISH: // client <-> server
		//fmt.Println("[trace] create publish message object")
		//pub := newPublishMsg()
		//pub.fixedHeader = header
		//msg = pub
	case PUBACK:
	case PUBREC: // server <-> client
	case PUBREL: // client <-> server
	case PUBCOMP:
	case SUBSCRIBE: // client -> server
		//fmt.Println("[trace] create subscribe message object")
		//sub := NewSubscribeMsg()
		//sub.fixedHeader = header
		//msg = sub
	case SUBACK:
	case UNSUBSCRIBE:
	case UNSUBACK:
	case PINGREQ: // client -> server
		//fmt.Println("[trace] create pingreq message object")
		//msg = NewPingreqMsg()
	case PINGRESP: // server -> client
		//fmt.Println("[trace] create pingresp message object")
	case DISCONNECT: // client -> server
		//fmt.Println("[trace] create disconnect message object")
	case RESERVED15:
	}
	msg.decode(b[l:])
	return msg
}

type fixedHeader struct {
	msgType msgType
	dup     bool
	qos     byte
	retain  bool
	remain  uint32
}

func decodeFixedHeader(b []byte) (*fixedHeader, int) {
	remain, l := decodeRemain(b[1:])
	return &fixedHeader{
		msgType: msgType(b[0] >> 4),
		dup:     b[0]&0x04 == 0x04,
		qos:     (b[0] & 0x06) >> 1,
		retain:  b[0]&0x01 == 0x01,
		remain:  remain,
	}, 1 + l
}
func (h *fixedHeader) encode(l uint32) []byte {
	buff := make([]byte, 0, 5)

	b := byte(h.msgType << 4)
	if h.dup {
		b |= 0x08 // 0x01 << 3
	}
	b |= h.qos << 1
	if h.retain {
		b |= 0x01
	}

	buff = append(buff, b)
	return append(buff, encodeRemain(l)...)
}

func decodeRemain(b []byte) (uint32, int) {
	remain := uint32(0)
	l := 0
	for i := 0; i < 4; i++ {
		if b[0]>>7 == 1 {
			remain += uint32(b[0] & byte(0x7f))
		} else {
			remain += uint32(b[0]&byte(0x7f)) << uint(i*7)
			l = i + 1
			break
		}
	}
	return remain, l
}

func encodeRemain(remain uint32) []byte {
	b := make([]byte, 0, 4)
	r := remain
	for i := 0; i < 4; i++ {
		l := byte(r) & 0x7F
		r >>= 7
		if r > 0 {
			l |= 0x80
			b = append(b, l)
		} else {
			b = append(b, l)
      break
		}
	}
	return b
}

// decode UTF8 string with length.
// 0:LENGTH MSB
// 1:LENGTH LSB
// 2: string
// n: ...
func decodeUTF8(b []byte) (string, int) {
	l := int(b[0]<<8 + b[1])
	str := b[2 : l+2]
	return string(str), l + 2
}
