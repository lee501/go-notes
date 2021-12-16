package router

import "github.com/gin-gonic/gin"

func Initrouter() *gin.Engine {
	router := gin.New()
	router.Any("/cpu/use", api.CpuInfo)
	return router
}

