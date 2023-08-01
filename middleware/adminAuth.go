package middleware

import "github.com/gin-gonic/gin"

func AdminAuthentication(c *gin.Context) {
	c.Next()
}
