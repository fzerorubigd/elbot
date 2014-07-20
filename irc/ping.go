package irc

import (
	"log"
	"reflect"
)

type IrcPing struct {
}

//Is this app supports this request?
func (app IrcPing) IsSupported(m *Message) bool {
	return m.Command() == "PING"
}

// Process the message
func (app IrcPing) Process(m *Message, w Writer) {
	if err := w.Command("PONG", m.Trailing()); err != nil {
		panic(err)
	}
	log.Println("Pong")
}

func NewPing() *IrcPing {
	return &IrcPing{}
}

func init() {
	RegisterAppType("ping", reflect.TypeOf(IrcPing{}))
}
