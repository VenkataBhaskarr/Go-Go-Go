package routes

import (
	"github.com/gin-gonic/gin"
	"gogogo/middleware"
)

func GetAdminRoutes(router *gin.Engine) *gin.RouterGroup {
	admins := router.Group("/admins")
	admins.POST("/signup", func(c *gin.Context) {
		// some random code
	})

	admins.POST("/login", func(c *gin.Context) {
		// some random code
	})

	admins.GET("/courses", middleware.AdminAuthentication, func(c *gin.Context) {
		// some random code
	})

	admins.POST("/courses/:courseId", middleware.AdminAuthentication, func(c *gin.Context) {

	})

	admins.GET("/purchasedCourses", middleware.AdminAuthentication, func(c *gin.Context) {

	})
	return admins
}
