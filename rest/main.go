package main

import (
	"github.com/Vladimir77715/adService/core/database"
	"github.com/Vladimir77715/adService/core/health"
	"github.com/Vladimir77715/adService/rest/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func PanicToResponseError(c *gin.Context) {
	defer func() {
		r := recover()
		if r != nil {
			c.Status(http.StatusInternalServerError)
			c.Writer.WriteString("Unexpected exception")
			c.Abort()
		}
	}()
	c.Next()
}

func initDbConnection() error {
	database.Client = &database.PostgresClient{Host: "localhost", Port: "5400",
		DbName: "addb", User: "ad_super_user", Password: "mxBCNjSqcw77", SslMode: database.NoSSl}
	return database.Client.IitDbConn()
}
func initGIN() chan error {
	router := gin.Default()
	health.Init()
	router.Use(PanicToResponseError)
	api := router.Group("/api")
	api.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message":       "pong",
			"serviceUpTime": health.ServiceStartTime.Format("Mon Jan 2 15:04:05 MST 2006"),
		})
	})
	services.RegisterAdService(api)
	errChan := make(chan error)
	go func() {
		if err := router.Run(":8080"); err != nil {
			errChan <- err
		}
	}()
	return errChan
}

func main() {
	if err := initDbConnection(); err != nil {
		log.Fatalf("Fatal error : %v", err)
		return
	}
	defer database.Client.CloseDbConn()
	erChan := <-initGIN()
	log.Fatalf("Fatal error : %v", erChan)
}
