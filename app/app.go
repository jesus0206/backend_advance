package app

import (
	"fmt"
	"log"
	"os"
	"yofio/avanzado/config"
	"yofio/avanzado/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// StartApp function runner server
func StartApp() {
	router := RouterInitial()
	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}

// RouterInitial exported
func RouterInitial() *gin.Engine {
	dbSQL := config.SQLServer{
		URLBD:      os.Getenv("BD_URL"),
		NameBD:     os.Getenv("BD_NAME"),
		UserBD:     os.Getenv("BD_USER_NAME"),
		PasswordBD: os.Getenv("BD_USER_PASSWORD"),
		PortBD:     os.Getenv("BD_PORT"),
	}

	db, err := config.GetConnectionPostgres(&dbSQL)
	if err != nil {
		fmt.Println("Error en la conexion con la bd..")
		log.Fatal(err)
	}

	if os.Getenv("DEBUG") == "TRUE" {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}
	// router.Use(cors.Default())

	controller := controllers.NewController(db)
	router.POST("/credit-assignment", controller.CreditAssignment)
	router.GET("/static", controller.Static)
	return router
}
