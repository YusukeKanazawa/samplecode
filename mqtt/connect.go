package mqtt


// connect flags.
// this flag will encode to a byte data that part of variable header.
type connectFlag struct {
	UserNameFlag bool
	PasswordFlag bool
	WillRetain   bool
	WillQoS      byte
	WillFlag     bool
	CleanSession bool
}

// bit structure of connect flags byte.
// bit | 7 | 6 | 5 | 4   3 | 2 | 1 | 0 |
//     |   |   |   |       |   |   | X |
//     |   |   |   |       |   | clean session
//     |   |   |   |       | will
//     |   |   |   | will QoS
//     |   |   | will Retain
//     |   | password
//     | user name
//
func newConnectFlag(b byte) *connectFlag {
	connectFlag := &connectFlag{
		UserNameFlag: b&0x80 == 0x80,
		PasswordFlag: b&0x40 == 0x40,
		WillRetain:   b&0x20 == 0x20,
		WillQoS:      b & 0x18 >> 3,
		WillFlag:     b&0x04 == 0x04,
		CleanSession: b&0x02 == 0x02,
	}
	return connectFlag
}

// header
//   Protocol: Only "MQIspd" ver "3" as MQTTv3.1 
//             or "MQTT" ver "4" as MQTTv3.1.1 supported.
//   keepAlive: 10 minutes default.
//
// payload
//   client id
//   will topic
//   will message
//   will username
//   username
//   password
type connectMsg struct {
  *fixedHeader
	Protocol    string
	Version     byte
	ConnectFlag *connectFlag
	KeepAlive   int
	ClientId    string
	WillTopic   string
	WillMessage string
	UserName    string
	Password    string
}

func newConnectMsg() *connectMsg {
  return &connectMsg{}
}

func (m *connectMsg) decode(b []byte) {
	cur := 0

	// variable header
  l:=0	
	m.Protocol, l = decodeUTF8(b[cur:])
	cur += l

	m.Version = b[cur]
	cur++

	m.ConnectFlag = newConnectFlag(b[cur])
	cur++

	m.KeepAlive = int(b[cur]<<8 | b[cur+1])
	cur += 2

	// payload
	/*
	 * under 23 chars
	 */
	m.ClientId, l = decodeUTF8(b[cur:])
	cur += l

	// TODO: length uder 23 check.
	if m.ConnectFlag.WillFlag {
		topic, l := decodeUTF8(b[cur:])
		cur += l

		message, l := decodeUTF8(b[cur:])
		cur += l

		// TODO: ASCII only check.
		m.WillTopic = topic
		m.WillMessage = message
	}
	if m.ConnectFlag.UserNameFlag {
		user, l := decodeUTF8(b[cur:])
		cur += l
		// TOOD: 0 <= length <= 12 (Warning)
		m.UserName = user
	}
	if m.ConnectFlag.PasswordFlag {
		password, l := decodeUTF8(b[cur:])
		cur += l
		m.Password = password
	}
}

func (m *connectMsg)encode()[]byte{
  return make([]byte,0)
}
