package mqtt

import ("log")


type broker struct {
	rcvCh chan message
}

func NewBroker() *broker {
	return &broker{
		rcvCh: make(chan message, 100),
	}
}
func (b *broker) GetRcvCh() chan message {
	return b.rcvCh
}

// go routine
// 
func (b *broker) WireMessage() {
  for{
    msg := <- b.rcvCh
		log.Println("test")
  }
}
