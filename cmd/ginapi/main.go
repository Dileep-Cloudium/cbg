package main

import (
	"context"
	"encoding/json"
	"log"

	routing "github.com/cbglabs/nexusgate-pbmapi/internal/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambdaV2

// Handler is main entry point for API proxy request coming to Lambda.
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	outCtx, err := json.Marshal(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Ctx: %s", string(outCtx))

	outReq, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	log.Printf("Req: %s", string(outReq))

	if !initialized {
		ginEngine := routing.MountRoutes()
		ginLambda = ginadapter.NewV2(ginEngine)
		initialized = true
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

// @title nexusgate-pbmapi
// @version 1.0
// @description This is for the NexusGate PBM API.
// @description This solution is intended to act as a collection of CBG services that support the CBG PBM APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	lambda.Start(Handler)
}
