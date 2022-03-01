package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting...")

	app := app.Build()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
