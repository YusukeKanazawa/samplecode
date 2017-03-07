package mqtt

import (
	"log"
	"io"
)
const(
	KEEP_ALIVE = 120
)
type session struct {
	Connected   bool
	Protocol    string
	Version     byte
	KeepAlive   int
	ClientId    string
	WillTopic   string
	WillMessage string
	UserName    string
	Rcv chan message
	Snd chan message
}

func NewSession() *session {
	return &session{
		Connected: false,
		KeepAlive: KEEP_ALIVE,
		Rcv: make(chan message, 10),
		Snd: make(chan message, 10),
	}
}

func (s *session) Receive() {
	for {
		msg := <-s.Rcv
		switch msg.getType(){
		case RESERVED0:
		case CONNECT: // client -> server
			 conMsg, _ := msg.(*connectMsg)
			 log.Println(s.ClientId + ":<< CONNECT")
			 s.handleConnect(conMsg)
			 log.Println(s.ClientId + ":>> CONNACK")
		case CONNACK: // server -> client
		case PUBLISH: // client <-> server
			log.Println(s.ClientId + ":<< PUBLISH")
			// pubMsg := m.(*publishMsg)
			// handlePublish(pubMsg)
			//				pub, _ := msg.(*PublishMessage)
			// qos implements...
			//go publish(pub.topic, pub)
			log.Println(s.ClientId + ":>> PUBACK")
			continue
		case PUBACK:
		case PUBREC: // server <-> client
		case PUBREL: // client <-> server
		case PUBCOMP:
		case SUBSCRIBE: // client -> server
			// log.Println("<< SUBSCRIBE")
			// subMsg, _ := m.(*subscribeMsg)
			// ackMsg, err := handleSubscribe(subMsg, ss)
			// if err != nil {
			// 	logger.Err(err.Error())
			// 	continue
			// }
			// ret = ackMsg
			// log.Println(">> SUBACK")
		case SUBACK:
		case UNSUBSCRIBE:
		case UNSUBACK:
		case PINGREQ: // client -> server
			log.Println(s.ClientId + ":<< PINGREQ")
			s.Snd <- newPingrespMsg()
			log.Println(s.ClientId + ":>> PINGRESP")
		case PINGRESP: // server -> client
		case DISCONNECT: // client -> server
			// ret = msg
		case RESERVED15:
		}

	}
}

func  (s *session) handleConnect(msg *connectMsg) {
	// protocol validate
	if msg.Protocol != "MQTT" || msg.Version != 4 {
		s.Snd <- newConnackMsg(CONNACK_PROTOOCL, false)
	}

	// clientId length validate
	// 1 <= length <= 23
	if idLen := len(msg.ClientId); idLen < 1 || idLen > 23 {
		s.Snd <- newConnackMsg(CONNACK_ID, false)
	}
	// start session with client id
	//ss, err := ssMgr.startSession(msg.ClientId, sndCh)
	//if err != nil {
		//return NewConnackMessage(CONNACK_NOT_AVAILABLE), nil
	//}

	//ss.SetCleanSession(msg.ConnectFlag.CleanSession)
	// success connect.


	s.Protocol = msg.Protocol
	s.Version = msg.Version
	if msg.KeepAlive > 0 {
		s.KeepAlive = msg.KeepAlive
	}
	s.ClientId = msg.ClientId
	s.WillTopic = msg.WillMessage
	s.WillMessage = msg.WillMessage
	s.Connected = true
	s.Snd <- newConnackMsg(CONNACK_OK, msg.ConnectFlag.CleanSession)
}
func (s *session) Send(c io.Writer){
	for {
		buff := <- s.Snd
		c.Write(buff.encode())
	}
}
