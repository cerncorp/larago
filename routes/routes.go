package routes

import (
	"fmt"

	"github.com/cerncorp/larago/app/http/controllers"
	"github.com/cerncorp/larago/app/services"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("routes package initialized")
}

func LoadRoutes(r *gin.Engine) (*gin.Engine, error) {
	callableController := controllers.NewCallableController(services.NewCallableService())

	r.GET("/", callableController.TestGet)
	r.POST("/", callableController.TestPost)
	r.PUT("/:id", callableController.TestPut)
	r.GET("/callservice/:name/:age", callableController.TestCallService)

	r.GET("/videos", callableController.TestGetVideo)
	r.POST("/videos", callableController.TestPostVideo)

	return r, nil
}
