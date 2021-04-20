package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Controllers package initialized")
}
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "heeloo Woorllddd",
	})
}
