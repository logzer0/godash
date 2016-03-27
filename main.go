package main

import (
	"log"

	"github.com/go-ini/ini"
)

var Config *ini.File

func main() {
	// Define buttons and action function
	DashMacs["74:c2:46:84:f5:ee"] = NoAction

	// Kick it off!
	SnifferStart()
}

func NoAction() {
	log.Println("No action on click.")
}
