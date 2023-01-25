package main

import (
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
	name := os.Args[0]
	if i := strings.LastIndexAny(name, "/\\"); i > -1 {
		name = name[i+1:]
	}
	log.SetPrefix(name + ": ")
}

func main() {
	log.Println("working")
}
