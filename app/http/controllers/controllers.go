package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Controllers package initialized")
}
func TestGet(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "heeloo Woorllddd! ",
		"name":    name,
		"age":     age,
	})
}

func TestPost(c *gin.Context) {
	body, err := c.GetRawData()
	// value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(404, gin.H{
			"result":  "false",
			"message": "bad request",
		})
		return
	}

	c.JSON(200, gin.H{
		"result":  "true",
		"message": string(body),
	})
}

func TestPut(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "id: " + id,
	})
}
