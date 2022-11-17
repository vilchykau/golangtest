package drivers

import (
	"github.com/gin-gonic/gin"
	"github.com/vilchykau/golangtest/internal/comerror"
)

func UnpackDatabase(c *gin.Context) (Database, error) {
	var dbRaw, ok = c.Get("DB")
	if !ok {
		return nil, comerror.ErrDatabaseNotInit
	}
	return dbRaw.(Database), nil
}
