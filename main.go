package main

import (
	"log"
	"os"
)

func main() {
	err := ExecuteTemplate(os.Stdin, os.Stdout, os.Args[1:]...)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
