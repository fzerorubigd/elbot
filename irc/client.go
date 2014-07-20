package irc

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"sync"
)

type Client struct {
	conn      net.Conn
	reader    *bufio.Reader
	writer    *bufio.Writer
	writeLock *sync.RWMutex
}

func (irc Client) Command(cmd string, data string) error {
	irc.writeLock.Lock()
	defer irc.writeLock.Unlock()

	writer := textproto.NewWriter(irc.writer)
	command := cmd
	if data != "" {
		command = command + " " + data
	}
	command = command + "\r\n"

	err := writer.PrintfLine(command)

	return err
}

func (irc Client) Read() (*Message, error) {
	reader := textproto.NewReader(irc.reader)
	line, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}

	return NewMessage(line)
}

func (irc Client) Close() error {
	return irc.conn.Close()
}

func NewClient(host string, port int) (*Client, error) {
	irc := Client{}
	var err error
	server := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(server)
	irc.conn, err = net.Dial("tcp", server)

	if err != nil {
		return nil, err
	}

	irc.reader = bufio.NewReader(irc.conn)
	irc.writer = bufio.NewWriter(irc.conn)

	irc.writeLock = new(sync.RWMutex)

	return &irc, nil
}
