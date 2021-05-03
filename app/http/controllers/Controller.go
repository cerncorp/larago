package controllers

import (
	"fmt"
	"net/http"

	"github.com/cerncorp/larago/app/models"
	"github.com/cerncorp/larago/app/modules"
	"github.com/cerncorp/larago/app/services"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func init() {
	fmt.Println("Controllers package initialized")
}

type CallableController interface {
	TestGet(*gin.Context)
	TestPost(c *gin.Context)
	TestPut(c *gin.Context)
	TestCallService(c *gin.Context)
	TestGetVideo(c *gin.Context)
	TestPostVideo(c *gin.Context)
}

type controller struct {
	service services.CallableService
}

var validate *validator.Validate

func NewCallableController(srv services.CallableService) CallableController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", modules.ValidateCoolTitle)

	return &controller{
		service: srv,
	}
}

// func Default() CallableController {
// 	return &controller{
// 		service: services.NewCallableService(),
// 	}
// }

func (cc *controller) TestGet(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"message": "heeloo Woorllddd! ",
		"name":    name,
		"age":     age,
	})
}

func (cc *controller) TestPost(c *gin.Context) {
	body, err := c.GetRawData()
	// value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

func (cc *controller) TestPut(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "id: " + id,
	})
}

// Testing call a service
// GET /callservice/:name/:age
func (cc *controller) TestCallService(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	// callableService := services.CallableService{}
	resp, _ := cc.service.CallService(name, age)

	c.JSON(200, gin.H{
		"message": resp,
	})
}

func (cc *controller) Save(c *gin.Context) error {
	var video models.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	cc.service.Save(video)
	c.JSON(http.StatusOK, gin.H{"message": video})
	return nil
}

// POST /videos
func (cc *controller) TestPostVideo(c *gin.Context) {
	err := cc.Save(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// GET /videos
func (cc *controller) TestGetVideo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": cc.service.FindAll()})
}
