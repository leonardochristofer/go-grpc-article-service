package main

import (
	"article-service/cmd"
	"log"
	"os"

	"article-service/config"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if cmd.Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
