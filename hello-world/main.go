package main

import (
	"log"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")

	r := gin.Default()

	ginLambda = ginadapter.New(r)
}

func main() {
	startHandling(&PingPongService{})
}
