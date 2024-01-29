package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mistadave/gt-app/config"
	"github.com/mistadave/gt-app/models"
	"github.com/mistadave/gt-app/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	models.Init()
	server.Init()
}
