package irc

import (
	"log"
	"reflect"
	"strings"
)

type Launcher struct {
	sys *System
}

//Is this app supports this request?
func (app Launcher) IsSupported(m *Message) bool {
	return m.Command() == "PRIVMSG" && m.trailing[0] == '@'
}

// Process the message
func (app Launcher) Process(m *Message, w Writer) {
	m.Print()
	command := strings.SplitN(m.trailing, " ", 2)
	switch command[0] {
	case "@launch":
		if len(command) > 1 {
			param := command[1]
			if typ, ok := appMap[param]; ok {
				nApp := reflect.New(typ).Elem().Interface()
				if v, ok := nApp.(App); ok {
					app.sys.Run(param, v)

				} else {
					log.Println("Not supported type", nApp)
				}
			} else {
				log.Println("Not found type", param, appMap)
			}
		} else {
			log.Println("Launh fail! : ", command)
		}
	case "@kill":
		if len(command) > 1 {
			param := command[1]
			err := app.sys.Kill(param)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println("Kill fail! : ", command)
		}
	case "@ps":
		for n, _ := range app.sys.apps {
			u, _, _ := m.GetUser()
			w.Command("PRIVMSG "+u, n)
		}
	}
}

func NewLauncher(sys *System) *Launcher {
	return &Launcher{sys}
}

func init() {
	RegisterAppType("launcher", reflect.TypeOf(Launcher{}))
}
