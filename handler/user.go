package handler

import (
	"net/http"

	"github.com/EikoNakashima/sample-gin/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Register(c *gin.Context) {
	input := model.UserInput{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	payload, err := service.Create(&input)

	c.JSON(http.StatusOK, payload)
}
