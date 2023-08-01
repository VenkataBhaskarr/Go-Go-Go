package main

import (
	"github.com/gin-gonic/gin"
	"gogogo/routes"
)

func main() {
	router := gin.Default()
	_ = routes.GetUserRoutes(router)
	_ = routes.GetAdminRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
