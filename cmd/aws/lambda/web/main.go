package main

import (
	"context"
	"embed"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/Control-Alt-Repeat/control-alt-repeat/internal/web"
)

var templates embed.FS
var echoLambda *echoadapter.EchoLambda

func init() {
	log.Printf("Echo cold start")
	var err error
	app, err := web.Init()
	if err != nil {
		log.Fatalf("Failed to initialize Echo app: %v", err)
	}
	echoLambda = echoadapter.New(app)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
