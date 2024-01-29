package server

import "github.com/mistadave/gt-app/config"

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	router.Run(config.GetString("server.port"))
}
