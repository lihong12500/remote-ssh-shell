package main

import (
	"ssh/Server"
	"ssh/command"
	"ssh/conf"
	"ssh/log"

	"github.com/spf13/viper"
)

type Message struct {
	Server  *Server.Server
	Content string
	Err     error
}

//var wg sync.WaitGroup
var results = make(chan Message, 255)
var stopChan = make(chan int, 255)

func main() {
	var finishedNum int
	var failedNum int
	conf.InitConfig()
	for _, server := range Server.InitServer() {
		go execute(server, results)
	}

	totalTask := viper.GetInt("server.total")
	//处理返回消息
	for {
		select {
		case result := <-results:
			if result.Err != nil {
				failedNum++
				log.Errorf("IP:%s-user:%s-port:%s -> %s",
					result.Server.Ip, result.Server.User, result.Server.Port, result.Err.Error())
			} else {
				log.Infof("IP:%s-user:%s-port:%s-> context:%s",
					result.Server.Ip, result.Server.User, result.Server.Port, result.Content)
			}
			finishedNum++
			stopChan <- finishedNum
		case flag := <-stopChan:
			log.Infof("tasks total: %d, finished: %d, unfinished: %d", totalTask, finishedNum, totalTask-finishedNum)
			if flag == totalTask {
				log.Info("all tasks success and program exit")
				if failedNum > 0 {
					log.Warnf("failed %d", failedNum)
				}
				return
			}
		}

	}
}

func execute(server *Server.Server, result chan Message) {
	cli := command.NewCli(server)
	output, err := cli.Run() // 执行的命令
	result <- Message{
		Server:  server,
		Content: output,
		Err:     err,
	}
}
