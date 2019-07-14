package controller

import (
	"net/http"

	"github.com/bxmon/diapi-mock-server/model"
	"github.com/bxmon/diapi-mock-server/service"
	"github.com/gin-gonic/gin"
)

// Controller defines controller structure
type Controller struct {
	service *service.Service
}

// NewController create new controller instance
func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}

// PingHandler handles ping server action
func (c *Controller) PingHandler(cx *gin.Context) {
	cx.JSON(http.StatusOK, gin.H{"message": "Mock server is running."})
}

// AddUserHandler handles add new user action
func (c *Controller) AddUserHandler(cx *gin.Context) {
	reqUser := new(model.User)
	if err := cx.ShouldBind(reqUser); err != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot bind request body"})
		return
	}

	if err := c.service.AddNewUser(reqUser); err != nil {
		cx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot add new user"})
		return
	}

	cx.JSON(http.StatusOK, gin.H{"userDetails": reqUser})
}

// GetUsersHandler handles get users action
func (c *Controller) GetUsersHandler(cx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		cx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot get all users"})
		return
	}

	cx.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserByIDHandler handles get user by id action
func (c *Controller) GetUserByIDHandler(cx *gin.Context) {
	params := struct {
		ID int `uri:"userid" binding:"required"`
	}{}

	err := cx.ShouldBindUri(&params)
	if err != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot bind request params"})
		return
	}

	user, err := c.service.GetUserByID(params.ID)
	if err != nil {
		cx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot get user"})
		return
	}

	cx.JSON(http.StatusOK, gin.H{"userDetails": user})
}

// UpdateUserHandler handles update user action
func (c *Controller) UpdateUserHandler(cx *gin.Context) {
	params := struct {
		ID int `uri:"userid" binding:"required"`
	}{}

	err := cx.ShouldBindUri(&params)
	if err != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot bind request params"})
		return
	}

	reqUser := new(model.User)
	if err := cx.ShouldBind(reqUser); err != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot bind request body"})
		return
	}

	if err := c.service.UpdateUser(reqUser); err != nil {
		cx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot update user"})
		return
	}

	cx.JSON(http.StatusOK, gin.H{"userDetails": reqUser})
}

// DeleteUserByIDHandler handles delete user action
func (c *Controller) DeleteUserByIDHandler(cx *gin.Context) {
	params := struct {
		ID int `uri:"userid" binding:"required"`
	}{}

	err := cx.ShouldBindUri(&params)
	if err != nil {
		cx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot bind request params"})
		return
	}

	if err := c.service.DeleteUserByID(params.ID); err != nil {
		cx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot delete user"})
		return
	}

	cx.JSON(http.StatusOK, gin.H{"message": "Success"})
}
