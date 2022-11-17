package db

import (
	"fmt"

	"github.com/vilchykau/golangtest/internal/drivers"
)

func GetDatabase() drivers.Database {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "habrpguser", "pgpwd4habr", "habrdb")
	driver :=
		drivers.NewPgDriver(drivers.PgConfig{DnsString: psqlInfo})
	err := driver.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}

	return driver
}
