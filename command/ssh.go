package command

import (
	"fmt"
	"net"
	"ssh/Server"

	"golang.org/x/crypto/ssh"
)

type Cli struct {
	User       string
	Pwd        string
	Addr       string
	Shell      string
	LastResult string
	client     *ssh.Client
	session    *ssh.Session
}

func NewCli(server *Server.Server) ICli {
	return &Cli{
		User:  server.User,
		Pwd:   server.Pwd,
		Addr:  fmt.Sprintf("%s:%s", server.Ip, server.Port),
		Shell: server.Shell,
	}
}

func (c *Cli) Connect() (*Cli, error) {
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.User = c.User
	config.Auth = []ssh.AuthMethod{ssh.Password(c.Pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }
	client, err := ssh.Dial(`tcp`, c.Addr, config)
	if nil != err {
		return c, err
	}
	c.client = client
	return c, nil
}

func (c *Cli) Run() (string, error) {
	if c.client == nil {
		_, err := c.Connect()
		if err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return ``, err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(c.Shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}
