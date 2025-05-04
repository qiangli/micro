package main

import (
	"os"
	"log"

	"github.com/zyedidia/micro/v2/editor"
)

func main() {
	args := os.Args[1:]
	go editor.NewEditor(args)

	rc := <- editor.Exiting

	log.Println("Exiting with code:", rc)
	os.Exit(rc)
}
