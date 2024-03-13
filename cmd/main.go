package main

import (
	"log"
	"os"

	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/models"
	"github.com/onemariovalentino/fetch-webpage/src/pkg/command"
)

var Metadata map[string]*models.Metadata

func main() {
	if len(os.Args) > 1 {
		cli := command.New(&Metadata)
		cli.Run()
	} else {
		log.Fatal("invalid argument")
		os.Exit(1)
	}
}
