package bootstrap

import (
	"log"

	"github.com/artur0us/test-work-311019/drivers/pgsqldrv"
	"github.com/artur0us/test-work-311019/servers/httpserver"
)

func All() {
	log.Println("[i] Connecting to databases...")
	pgsqldrv.ConnectAll()
	log.Println("[i] OK")

	log.Println("[i] Starting HTTP server...")
	httpserver.Start()
}
