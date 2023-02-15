package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}

type ApiGatewayService interface {
	Routes() []Route
}

type PingPongService struct{}

// Compile time interface check
var _ ApiGatewayService = (*PingPongService)(nil)

func (*PingPongService) Routes() []Route {
	return []Route{
		{Method: http.MethodGet, Path: "/ping", Handler: HandlePing},
	}
}

func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
