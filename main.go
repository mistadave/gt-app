package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mistadave/gt-api/config"
	"github.com/mistadave/gt-api/db"
	"github.com/mistadave/gt-api/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
