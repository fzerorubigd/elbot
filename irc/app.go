package irc

type Writer interface {
	Command(cmd string, data string) error
}

type App interface {

	//Is this app supports this request?
	IsSupported(m *Message) bool

	// Process the message
	Process(m *Message, w Writer)
}
