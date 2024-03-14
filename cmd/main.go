package main

import (
	"log"
	"os"

	"github.com/onemariovalentino/fetch-webpage/src/pkg/command"
)

func main() {
	if len(os.Args) > 1 {
		cli := command.New()
		cli.Run()
	} else {
		log.Fatal("invalid argument")
		os.Exit(1)
	}
}
