package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func RequireQueryIntDefault(c *gin.Context, query string, defValue int) (int, error) {
	v := c.Query(query)
	if v == "" {
		return defValue, nil
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}
