package main

import (
	"github.com/aveplen-bach/config-gateway/internal/config"
	"github.com/aveplen-bach/config-gateway/internal/server"
)

func main() {
	var cfg config.Config
	config.ReadConfig("config-gateway.yaml", &cfg)
	server.Start(cfg)
}
