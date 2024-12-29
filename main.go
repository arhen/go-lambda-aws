package main

import (
	"fmt"
	"net/http"
    "context"

    "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var lambdaHandler *httpadapter.HandlerAdapter

// Handle requests for the root path
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Lambda!")
}

// Handle requests for the /api/ping path
func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

// Handle requests for the /api/status path
func handleStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status":"ok","message":"API is running smoothly"}`)
}

// Handle requests for the /api/hello/{name} path
func handleHello(w http.ResponseWriter, r *http.Request) {
	// Extract name from URL (URL params parsing needs to be done manually)
	// Here we'll just use a simple way to extract the path part after `/hello/`
	name := r.URL.Path[len("/api/hello/"):]
	if name == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func init() {
	// Define routes and handlers
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/ping", handlePing)
	http.HandleFunc("/api/status", handleStatus)
	http.HandleFunc("/api/hello/", handleHello)

	// Wrap the HTTP handler to be Lambda-compatible
	lambdaHandler = httpadapter.New(http.DefaultServeMux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
    return lambdaHandler.ProxyWithContext(ctx, req)
}


func main() {
    lambda.Start(Handler)
}   
