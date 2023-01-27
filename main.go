package main

import (
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
	name := os.Args[0]
	log.SetPrefix(name[strings.LastIndexAny(name, "/\\")+1:] + ": ")
}

func main() {
	log.Println("working")
}
