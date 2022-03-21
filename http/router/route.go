package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sekou-diarra/fiz-buzz-server/http/handler"
	"io"
	"time"
)

func InitRouter(fbHandler handler.FizzBuzzHandler) *gin.Engine {
	r := gin.New()

	ginLogger := logger.WithLogger(
		func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
			return log.With().
				Str("path", c.Request.URL.Path).
				Str("method", c.Request.Method).
				Dur("latency", latency).
				Logger()
		},
	)

	r.Use(cors.Default())
	r.Use(logger.SetLogger(ginLogger))
	r.Use(gin.Recovery())
	r.NoRoute(fbHandler.Health())

	r.RedirectTrailingSlash = true

	r.GET("/api/compute/fizzbuzz", fbHandler.GetFizzBuzz())
	r.GET("/api/metrics/besthits", fbHandler.GetBestHits())
	r.GET("/health", fbHandler.Health())

	return r
}
