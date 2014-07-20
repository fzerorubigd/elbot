package irc

import (
	"fmt"
	"log"
	"reflect"
)

var appMap map[string]reflect.Type = make(map[string]reflect.Type)

type System struct {
	irc  *Client
	apps map[string]App
}

func (sys *System) Run(name string, app App) error {
	if _, ok := sys.apps[name]; ok {
		return fmt.Errorf("Application with name %s is in run state", name)
	}
	sys.apps[name] = app

	return nil
}

func (sys *System) Kill(name string) error {
	if _, ok := sys.apps[name]; !ok {
		return fmt.Errorf("Application with name %s is in not ruunning", name)
	}
	delete(sys.apps, name)
	fmt.Println(sys.apps)
	return nil
}

func (sys *System) processLine(line *Message, name string, app App) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Killing %s due to panic. recover data : %s ", name, r)
			if _, ok := sys.apps[name]; ok {
				delete(sys.apps, name)
			}
		}
	}()
	if app.IsSupported(line) {
		log.Println(line.command, " Accepted by ", name)
		app.Process(line, sys.irc)
	}
}

func (sys *System) Process(nick string, user string, real string) error {

	//This part is essential. its better to be here
	sys.irc.Command("NICK", nick)
	sys.irc.Command("USER "+user+" 0 * :", real)

	for {
		line, err := sys.irc.Read()

		if err != nil {
			return err
		}
		//fmt.Println(line.Command())

		for name, app := range sys.apps {
			go sys.processLine(line, name, app)
		}
	}
}

func NewSystem(irc *Client) *System {
	sys := System{irc, make(map[string]App)}

	return &sys
}

func RegisterAppType(name string, typ reflect.Type) {
	appMap[name] = typ
}
