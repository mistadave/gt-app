package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mistadave/gt-app/config"
	"github.com/mistadave/gt-app/internal/models"
	"github.com/mistadave/gt-app/internal/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	models.DBInit()
	server.Init()
}
