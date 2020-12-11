package handler

import (
	"fiz-buzz-server/model"
	fb "fiz-buzz-server/service/fizzbuzz"
	"fiz-buzz-server/service/stats"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type FizzBuzzHandler interface {
	GetFizzBuzz() gin.HandlerFunc
	GetBestHits() gin.HandlerFunc
	Health() gin.HandlerFunc
}

type fizzBuzzHandler struct {
	fbService   fb.FizzBuzzService
	statService stats.StatService
}

func NewFizzBuzzHandler(fbservice fb.FizzBuzzService, statService stats.StatService) FizzBuzzHandler {
	return &fizzBuzzHandler{fbService: fbservice, statService: statService}
}

func (f fizzBuzzHandler) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	}
}

func (f fizzBuzzHandler) GetFizzBuzz() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params = model.FizzbuzzParam{}
		err := c.ShouldBindQuery(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var result = f.fbService.ComputeFizzBuzz(params)
		f.statService.SaveQuery(params)
		c.JSON(http.StatusOK, gin.H{"result": strings.Join(result.Sequences, ",")})
	}
}

func (f fizzBuzzHandler) GetBestHits() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := f.statService.GetMostAskedQuery()
		if r == nil {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusOK, r)
	}
}
