package constants

import (
	"../mdls"
)

func GetAllConnConfigs() []mdls.ConnConfig {
	return []mdls.ConnConfig{
		mdls.ConnConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "artur0us",
			Password: "xxxxxxxxxx",
			DBName:   "developer_notes",
		},
	}
}
