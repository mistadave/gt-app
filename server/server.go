package server

import "github.com/mistadave/gt-api/config"

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	router.Run(config.GetString("server.port"))
}
