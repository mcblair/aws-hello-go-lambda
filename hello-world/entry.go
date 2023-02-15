package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

const (
	// LambdaPortEnvVar defines the environment variable where the port is stored in a lambda environment.
	LambdaPortEnvVar = "_LAMBDA_SERVER_PORT"

	// APIRuntimeEnvVar defines the environment variable where the API runtime is stored in a lambda environment.
	APIRuntimeEnvVar = "AWS_LAMBDA_RUNTIME_API"
)

func startHandling(service ApiGatewayService) {
	if isLambdaEnv() {
		lambda.Start(ginHandler)
	} else {
		localStart(service.Routes())
	}
}

func ginHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func localStart(routes []Route) {
	r := gin.Default()

	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	r.Run(":8080")
}

func isLambdaEnv() bool {
	return len(os.Getenv(LambdaPortEnvVar)) > 0 || len(os.Getenv(APIRuntimeEnvVar)) > 0
}
