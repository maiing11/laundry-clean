package config

import "github.com/gin-gonic/gin"

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler() *RequestHandler {
	engine := RequestHandler{Gin: gin.Default()}
	return &engine
}
