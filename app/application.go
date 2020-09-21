package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("starting application on port 8080")
	fmt.Println(router.Run(":8080"))
}
