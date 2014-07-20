package irc

// Base of this code is from https://github.com/husio/go-irc

import (
	"errors"
	"log"
	"strings"
)

var (
	UnknownCommandError = errors.New("Unknown command")
)

type Message struct {
	raw      string
	prefix   string
	command  string
	params   []string
	trailing string
}

func (m *Message) Command() string {
	return m.command
}

func (m *Message) Trailing() string {
	return m.trailing
}

func (m *Message) Prefix() string {
	return m.prefix
}

func (m *Message) Params() []string {
	return m.params
}

func (m *Message) GetUser() (nick string, ident string, host string) {
	parts := strings.SplitN(m.prefix, "!", 2)
	if len(parts) == 2 {
		nick = parts[0]
		iParts := strings.SplitN(parts[1], "@", 2)
		if len(iParts) == 2 {
			ident = iParts[0]
			host = iParts[1]
		}
	}

	return
}

func (m *Message) Print() {
	log.Println("------")
	log.Println(m.raw)
	log.Println(m.prefix)
	log.Println(m.command)
	log.Println(m.params)
	log.Println(m.trailing)
	u, n, h := m.GetUser()
	log.Println(u, "-", n, "-", h)
	log.Println("------")
}

func NewMessage(raw string) (*Message, error) {
	raw = strings.TrimSpace(raw)
	msg := &Message{raw: raw}
	if raw[0] == ':' {
		chunks := strings.SplitN(raw, " ", 2)
		msg.prefix = chunks[0][1:]
		raw = chunks[1]
	}
	chunks := strings.SplitN(raw, " ", 2)
	msg.command = chunks[0]
	raw = chunks[1]
	if msg.command == "" {
		return nil, UnknownCommandError
	}

	if raw[0] != ':' {
		chunks := strings.SplitN(raw, " :", 2)
		msg.params = strings.Split(chunks[0], " ")
		if len(chunks) == 2 {
			raw = chunks[1]
		} else {
			raw = ""
		}
	}

	if len(raw) > 0 {
		if raw[0] == ':' {
			raw = raw[1:]
		}
		msg.trailing = raw
	}
	return msg, nil
}
