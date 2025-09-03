package handlers

import (
	"FitByte/configs"
	"FitByte/internal/middleware"
	"FitByte/internal/service"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	Engine    *gin.Engine
	AppConfig configs.Config
	FileSvc   service.FileService
}

func NewFileHandler(engine *gin.Engine, appConfig configs.Config, fileService service.FileService) *FileHandler {
	return &FileHandler{
		Engine:    engine,
		AppConfig: appConfig,
		FileSvc:   fileService,
	}
}

func (h *FileHandler) SetupRoutes() {
	routes := h.Engine.Group("/file")
	routes.Use(middleware.RequestLogger())
	routes.Use(middleware.AuthMiddleware(h.AppConfig.Secret.JWTSecret))
	routes.POST("upload", h.Upload)
}

func (h *FileHandler) Upload(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "upload endpoint",
	})
}
