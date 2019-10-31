package pgsqldrv

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"./mdls"
)

var All map[string]*sql.DB

func ConnectAll() {
	All = make(map[string]*sql.DB)

	allConnConfigs := []mdls.ConnConfig{
		mdls.ConnConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "artur0us",
			Password: "xxxxxxxxxx",
			DBName:   "developer_notes",
		},
	}

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
