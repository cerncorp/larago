package routes

import (
	"fmt"

	"github.com/cerncorp/larago/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("routes package initialized")
}

func LoadRoutes() (*gin.Engine, error) {
	r := gin.Default()
	r.GET("/", controllers.TestGet)
	r.POST("/", controllers.TestPost)
	r.PUT("/:id", controllers.TestPut)

	return r, nil
}
