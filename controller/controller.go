package controller

import (
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

// AddUserHandler handles add new user action
func (c *Controller) AddUserHandler(cx *gin.Context) {
}

// GetUsersHandler handles get users action
func (c *Controller) GetUsersHandler(cx *gin.Context) {
}

// GetUserByIDHandler handles get user by id action
func (c *Controller) GetUserByIDHandler(cx *gin.Context) {
}

// UpdateUserHandler handles update user action
func (c *Controller) UpdateUserHandler(cx *gin.Context) {
}

// ReplaceUserHandler handles update user action
func (c *Controller) ReplaceUserHandler(cx *gin.Context) {
}

// DeleteUserByIDHandler handles delete user action
func (c *Controller) DeleteUserByIDHandler(cx *gin.Context) {
}
