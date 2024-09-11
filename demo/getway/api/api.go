package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Usage float64
}

func CpuInfo(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Usage: 100.0,
	})
}
