package routes

import "github.com/gin-gonic/gin"
import "gogogo/middleware"

func GetUserRoutes(router *gin.Engine) *gin.RouterGroup {
	// some random code
	users := router.Group("/users")
	users.POST("/signup", func(c *gin.Context) {
		// some random code
	})

	users.POST("/login", func(c *gin.Context) {
		// some random code
	})

	users.GET("/courses", middleware.UserAuthentication, func(c *gin.Context) {
		// some random code
	})

	users.POST("/courses/:courseId", middleware.UserAuthentication, func(c *gin.Context) {

	})

	users.GET("/purchasedCourses", middleware.UserAuthentication, func(c *gin.Context) {

	})

	return users
}
