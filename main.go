package main

import (
	"fiz-buzz-server/http/handler"
	"fiz-buzz-server/http/router"
	log "fiz-buzz-server/logging"
	"fiz-buzz-server/repository/in_memory"
	fbservice "fiz-buzz-server/service/fizzbuzz"
	fbstat "fiz-buzz-server/service/stats"

	//statService "fiz-buzz-server/service/stats"
	"fmt"
	"github.com/gin-gonic/gin"
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
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.CustomLog.Fatal().Err(err).Msgf("error while starting the app: \n %s", err.Error())
	}
}
