package main

import (
	"log"
	"os"
	"term/internal/ter"
)

func main() {
	argsFile := os.Args[1]

	file, err := os.OpenFile(argsFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	editor := ter.NewEditor(file)
	editor.Start()
}
