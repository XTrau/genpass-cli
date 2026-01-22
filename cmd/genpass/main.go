package main

import (
	"genpass/external/app"
	"log"
	"os"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
