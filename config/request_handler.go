package config

import "github.com/gin-gonic/gin"

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler() RequestHandler {
	engine := gin.Default()
	return RequestHandler{Gin: engine}
}
