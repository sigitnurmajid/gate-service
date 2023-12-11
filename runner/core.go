package runner

import (
	"context"
	"fmt"
	"gate-service/config"
	"gate-service/service"
	"gate-service/zoo"
	"log"
)

const LOGO = `
    _  _____ _____ ____  ___
   / \|_   _| ____|  _ \|_ _|
  / _ \ | | |  _| | |_) || |
 / ___ \| | | |___|  _ < | |
/_/   \_\_| |_____|_| \_\___|
`

const SERVICENAME = "Gate Service"
const VERSION = "v0.1.0"

func Run(ctx context.Context, configPath string) {
	fmt.Print(LOGO + "\n" + SERVICENAME + " " + VERSION + "\n\n")

	conf, err := config.LoadFromFile(configPath)
	if err != nil {
		panic(err)
	}

	log.Println("Build TCP Server")
	server := service.NewTcpServer(conf.Port)

	log.Println("Create Zoo API")
	api := zoo.CreateZooApi(conf.ZooApi)

	log.Println("Create Session Memory")
	session := NewSessionStore()

	go func() {
		for msg := range server.Msgch {
			event := CreateEvent(api, msg, session)
			go event.Run()
		}
	}()

	log.Fatal(server.Start())
}
