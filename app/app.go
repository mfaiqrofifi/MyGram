package app

import (
	"MyGram/config"
	"MyGram/repository"
	"MyGram/route"
	"MyGram/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)
	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
