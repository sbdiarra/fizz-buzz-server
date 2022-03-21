package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sekou-diarra/fiz-buzz-server/http/handler"
	"github.com/sekou-diarra/fiz-buzz-server/http/router"
	"github.com/sekou-diarra/fiz-buzz-server/repository/in_memory"
	fbservice "github.com/sekou-diarra/fiz-buzz-server/service/fizzbuzz"
	fbstat "github.com/sekou-diarra/fiz-buzz-server/service/stats"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/sekou-diarra/fiz-buzz-server/logging"
)

func initApp() *gin.Engine {

	var inMemoryStore = in_memory.NewInMemoryStore()
	var statService = fbstat.NewStatService(inMemoryStore)
	var fbService = fbservice.NewFizzBuzzService()
	var fbHandler = handler.NewFizzBuzzHandler(fbService, statService)
	var r = router.InitRouter(fbHandler)

	return r

}

func main() {
	r := initApp()
	log.Info().Msg("initiating app")
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal().Err(err).Msg("error while starting the app")
	}
	log.Info().Msg("server started")
}
