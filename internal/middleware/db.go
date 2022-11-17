package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vilchykau/golangtest/internal/drivers"
)

func Db() gin.HandlerFunc {
	return func(c *gin.Context) {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"127.0.0.1", 5432, "habrpguser", "pgpwd4habr", "habrdb")
		driver :=
			drivers.NewPgDriver(drivers.PgConfig{DnsString: psqlInfo})
		err := driver.Connect()
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Set("DB", driver)
		c.Next()
	}
}
