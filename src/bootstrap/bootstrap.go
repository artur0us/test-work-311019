package bootstrap

import (
	"log"

	"../drivers/pgsqldrv"
	"../servers/httpserver"
)

func All() {
	log.Println("[i] Connecting to databases...")
	pgsqldrv.ConnectAll()
	log.Println("[i] OK")

	log.Println("[i] Starting HTTP server...")
	httpserver.Start()
}
