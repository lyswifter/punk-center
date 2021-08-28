package main

import (
	"log"
)

const repoPath = "~/.punkcenter"

func main() {
	log.Print("This is the punk center server")

	DataStores()

	SetupServers()
}
