package handlers

import (
	"FitByte/internal/models"
	"FitByte/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	customErrors "FitByte/internal/errors"
)

type UserHandler struct {
	Engine  *gin.Engine
	UserSvc service.UserService
}

func NewUserHandler(engine *gin.Engine, userService service.UserService) *UserHandler {
	return &UserHandler{
		Engine:  engine,
		UserSvc: userService,
	}
}

func (h *UserHandler) SetupRoutes() {
	routes := h.Engine.Group("/user")
	routes.POST("register", h.Register)
	routes.GET("health-check", h.pong)
}

func (h *UserHandler) pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *UserHandler) Register(c *gin.Context) {
	var model models.User
	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.UserSvc.Register(ctx, model)
	if err != nil {
		if errors.Is(customErrors.ErrUserAlreadyExists, err) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
