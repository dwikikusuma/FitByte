package handlers

import (
	"FitByte/configs"
	"FitByte/internal/middleware"
	"FitByte/internal/models"
	"FitByte/internal/service"
	"FitByte/pkg/log"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

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
	routes := h.Engine.Group("/v1/file")
	routes.Use(middleware.AuthMiddleware(h.AppConfig.Secret.JWTSecret))
	routes.POST("", h.Upload)
}

func (h *FileHandler) Upload(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt64("user_id")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to get file from request",
		})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Logger.Error().Err(err).Msg("Failed to close file")
		}
	}(file)

	timestamp := time.Now().Unix()
	uploadedFileName := header.Filename
	ext := filepath.Ext(uploadedFileName)
	filename := fmt.Sprintf("%d_%d%s", userID, timestamp, ext)

	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	fileModel := models.UploadFile{
		FileName:    filename,
		FileData:    file,
		Size:        header.Size,
		ContentType: contentType,
		FilePath:    fmt.Sprintf("/uploads/%d/%s", userID, filename),
	}

	uri, err := h.FileSvc.SaveFileUpload(ctx, userID, fileModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to upload file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uri": uri,
	})
}
