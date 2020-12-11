package router

import (
	"fiz-buzz-server/http/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(fbHandler handler.FizzBuzzHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.RedirectTrailingSlash = true

	r.GET("/api/compute/fizzbuzz", fbHandler.GetFizzBuzz())
	r.GET("/api/metrics/besthits", fbHandler.GetBestHits())
	r.GET("/", fbHandler.Health()).GET("/health", fbHandler.Health())

	return r
}
