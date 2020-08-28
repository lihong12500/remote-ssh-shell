package Server

import (
	"ssh/log"
	"strconv"

	"github.com/spf13/viper"
)

type Server struct {
	User  string
	Pwd   string
	Ip    string
	Port  string
	Shell string
}

var tasks = make(map[string]*Server, 0)

func InitServer() map[string]*Server {
	var total = viper.GetInt("server.total")

	//range all host information to add map
	for i := 1; i < total+1; i++ {
		key := "server" + strconv.Itoa(i)
		server := getValue(key)
		log.Debugf("add %s to tasks %s", key, server)
		tasks[key] = server
	}
	return tasks
}

func getValue(key string) *Server {
	user := viper.GetString(key + ".user")
	if user == "" {
		user = viper.GetString("server.user")
	}
	pwd := viper.GetString(key + ".pwd")
	if pwd == "" {
		pwd = viper.GetString("server.pwd")
	}
	ip := viper.GetString(key + ".ip")
	if ip == "" {
		ip = viper.GetString("server.ip")
	}
	shell := viper.GetString(key + ".shell")
	if shell == "" {
		shell = viper.GetString("server.shell")
	}
	port := viper.GetString(key + ".port")
	if port == "" {
		port = viper.GetString("server.port")
	}

	return &Server{
		User:  user,
		Pwd:   pwd,
		Ip:    ip,
		Port:  port,
		Shell: shell,
	}
}
