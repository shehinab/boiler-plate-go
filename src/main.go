package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/abhishek.k8/crud/src/config"
	"gitlab.com/abhishek.k8/crud/src/cron"
	"gitlab.com/abhishek.k8/crud/src/database"
	"gitlab.com/abhishek.k8/crud/src/migration"
	route "gitlab.com/abhishek.k8/crud/src/routes"
)

var router *gin.Engine

func main() {
	//initialize application with toml file
	if err := config.Init(); err != nil {
		log.Error(err)
	}
	// router = gin.Default()
	router = gin.New()
	if strings.ToLower(config.AppConfig.Environment) == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	mainRouter := new(route.MainRouter)
	appVer := "api/v1"
	mainRouter.GetRoutes(router.Group(appVer))
	//start the databse
	dbconn := database.ConnectSQL()
	defer dbconn.Close()
	migration.Migrate()

	//initialting cron
	var cron = cron.Cron{}
	cron.Init()

	//server starting log
	log.Info("Server starting at :: ", config.AppConfig.Server.Host+":"+config.AppConfig.Server.Port)

	router.Run(":" + config.AppConfig.Server.Port)
}

//CORSMiddleware -
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.BindHeader()
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Content-Length", "402653184")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization,Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST , HEAD , PATCH , OPTIONS, GET, PUT, DELETE")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Content-Length", "402653184")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,authorization, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST , HEAD, PATCH , OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
