package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EikoNakashima/sample-gin/model"
	"github.com/EikoNakashima/sample-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Create(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.Create(&input)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetAll(c *gin.Context) {
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetOne(c *gin.Context) {
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetOne(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context) {
	fmt.Println("params:", c.Param("user-id"))
	// userID := c.Param("user-id")
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("userID:", userID)
	input := model.UserInput{}
	fmt.Println("input:", input)
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.Update(&input, userID)
	fmt.Println("payload:", payload)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Delete(c *gin.Context) {
	// userID := c.MustGet("user-id").(uint64)
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	err = user.Delete(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "削除されました")
}
