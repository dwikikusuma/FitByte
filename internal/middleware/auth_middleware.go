package middleware

import (
	"FitByte/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		partedHeader := strings.Split(authHeader, " ")
		if len(partedHeader) != 2 || partedHeader[0] != "Bearer" || partedHeader[1] == "" {
			log.Logger.Error().Msg("Unauthorized: Invalid Authorization header")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := partedHeader[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			log.Logger.Error().Err(err).Msg("Unauthorized: Invalid token")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Logger.Error().Msg("Unauthorized: Invalid token claims")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			log.Logger.Error().Msg("Unauthorized: user_id claim is invalid or missing")
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		currentTime := time.Now()
		c.Set("user_id", userID)

		currentUserID := c.GetFloat64("user_id")

		c.Next()

		latency := time.Since(currentTime)
		status := c.Writer.Status()
		logger := log.Logger.With().
			Float64("user_id", currentUserID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Str("client_ip", c.ClientIP()).
			Dur("latency", latency).
			Logger()

		if status >= 500 {
			logger.Error().Msg("request completed with server error")
		} else if status >= 400 {
			logger.Warn().Msg("request completed with client error")
		} else {
			logger.Info().Msg("request completed successfully")
		}
	}
}
