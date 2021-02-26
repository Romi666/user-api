package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-api/model"
	"user-api/user"
	"user-api/utils"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func CreatePersonHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	v1 := r.Group("/api/v1/")
	v1.POST("user", userHandler.create)
	v1.PATCH("user", userHandler.update)
	v1.GET("users", userHandler.getAll)
	v1.GET("user", userHandler.getByID)
}

func (e *UserHandler) create(c *gin.Context) {
	var user model.UserData

	err := c.BindJSON(&user)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Bad Request")
		return
	}

	err = e.userUsecase.Create(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, map[string]interface{}{"message" : "success create user"})
}

func (e *UserHandler) update(c *gin.Context) {
	var user model.UserData

	err := c.BindJSON(&user)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Bad Request")
		return
	}

	userID := c.Query("user_id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = e.userUsecase.Update(userIDInt, &user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, map[string]interface{}{"message" : "success update user"})
}

func (e *UserHandler) getAll(c *gin.Context) {
	userList, err := e.userUsecase.GetAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, userList)
}

func (e *UserHandler) getByID(c *gin.Context) {
	userID := c.Query("user_id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := e.userUsecase.GetByID(userIDInt)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, user)
}