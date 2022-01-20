package main

import (
	"github.com/Vladimir77715/adService/core/database"
	"github.com/Vladimir77715/adService/core/health"
	"github.com/Vladimir77715/adService/rest/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
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
func JSONMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()
}

func initDbConnection() error {
	database.Client = &database.PostgresClient{DbConn: os.Getenv("DB_CONN")}
	return database.Client.IitDbConn()
}
func initGIN() chan error {
	router := gin.Default()
	health.Init()
	router.Use(JSONMiddleware, PanicToResponseError)
	api := router.Group("/api")
	api.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message":       "pong",
			"serviceUpTime": health.GetServiceUptime(),
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
