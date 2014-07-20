package irc

import "reflect"

type Echo struct {
}

//Is this app supports this request?
func (app Echo) IsSupported(m *Message) bool {
	return m.Command() == "PRIVMSG"
}

// Process the message
func (app Echo) Process(m *Message, w Writer) {
	user, _, _ := m.GetUser()
	if err := w.Command("PRIVMSG "+user, m.Trailing()); err != nil {
		panic(err)
	}
}

func NewEcho() *Echo {
	return &Echo{}
}

func init() {
	RegisterAppType("echo", reflect.TypeOf(Echo{}))
}
