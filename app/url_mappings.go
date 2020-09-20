package app

import (
	"github.com/KatherineEbel/bookstore_users-api/controllers/ping"
	"github.com/KatherineEbel/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:userId", users.Get)
	// router.GET("/users/search", controllers.Search)
	router.POST("/users", users.Insert)
}
