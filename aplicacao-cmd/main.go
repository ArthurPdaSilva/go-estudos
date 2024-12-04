package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal()
	}
}
