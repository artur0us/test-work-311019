package main

import (
	"log"

	"./bootstrap"
)

func main() {
	log.Println("[i] Starting *developer notes* service...")

	bootstrap.All()
}
