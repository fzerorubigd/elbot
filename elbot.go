package main

import (
	"encoding/json"
	elbirc "github.com/fzerorubigd/elbot/irc"
	"log"
	"os"
	"os/user"
)

type Config struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
	Username string `json:"username"`
	Realname string `json:"realname"`
}

func loadConfig() *Config {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := usr.HomeDir + "/.config/elbot/config.json"
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := &Config{}
	err = decoder.Decode(conf)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func main() {
	conf := loadConfig()

	irc, err := elbirc.NewClient(conf.Server, conf.Port)

	if err != nil {
		log.Fatal(err)
	}
	defer irc.Close()

	system := elbirc.NewSystem(irc)

	system.Run("ping", elbirc.NewPing())
	system.Run("launcher", elbirc.NewLauncher(system))

	system.Process(conf.Nick, conf.Username, conf.Realname)
}
