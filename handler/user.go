package handler

import (
	"net/http"

	"github.com/EikoNakashima/sample-gin/model"
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
	payload, err := service.Create(&input)

	c.JSON(http.StatusOK, payload)
}

func GetAll(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	payload, err := service.GetAll(&input)

	c.JSON(http.StatusOK, payload)
}

func GetOne(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	payload, err := service.GetAll(&input)

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	payload, err := service.Update(&input)

	c.JSON(http.StatusOK, payload)
}

func Delete(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	payload, err := service.Delete(&input)

	c.JSON(http.StatusOK, payload)
}
