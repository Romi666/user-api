package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-api/model"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	responseData := model.Response{
		Status: strconv.Itoa(http.StatusOK),
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, responseData)
}

func HandleError(c *gin.Context, status int, message string) {
	responseData := model.Response{
		Status: strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responseData)
}
