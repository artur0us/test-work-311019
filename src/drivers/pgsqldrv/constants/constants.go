package constants

import (
	"github.com/artur0us/test-work-311019/drivers/pgsqldrv/mdls"
)

func GetAllConnConfigs() []mdls.ConnConfig {
	return []mdls.ConnConfig{
		mdls.ConnConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "postgres",
			DBName:   "test_work",
		},
	}
}
