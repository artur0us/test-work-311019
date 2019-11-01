package pgsqldrv

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/artur0us/test-work-311019/drivers/pgsqldrv/constants"
	"github.com/artur0us/test-work-311019/drivers/pgsqldrv/mdls"
)

var All map[string]*sql.DB

func ConnectAll() {
	All = make(map[string]*sql.DB)

	allConnConfigs := constants.GetAllConnConfigs()

	for _, oneConnConfig := range allConnConfigs {
		if err := connect(oneConnConfig); err != nil {
			log.Println("[!] Critical error occurred while connecting to database! More info: ")
			panic(err)
		}
	}
}

func connect(connConfig mdls.ConnConfig) error {
	connStr := "host=" + connConfig.Host + " port=" + connConfig.Port + " user=" + connConfig.Username + " password=" + connConfig.Password + " dbname=" + connConfig.DBName + " sslmode=disable"
	var err error
	All[connConfig.DBName], err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	_, err = All[connConfig.DBName].Exec("SELECT 1")
	if err != nil {
		return err
	}
	return nil
}
