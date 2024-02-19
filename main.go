package main

import (
	"log"

	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/di"
)

func main() {

	cfg,cfgErr:=config.LoadEnvVariables()
	if cfgErr!=nil{
		log.Fatal("cannot load config: ",cfgErr)
	}

	server,diErr:=di.InitializeAPI(cfg)
	if diErr!=nil{
		log.Fatal("cannot start server: ",diErr)
	} else {
		server.Start()
	}
}