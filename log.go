package main

import (
	"log"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	//"golang.org/x/net/context"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "point-c"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	logName := "my-log"

	logger := client.Logger(logName).StandardLogger(logging.Info)

	// Logs "hello world", log entry is visible at
	// Stackdriver Logs.
	logger.Println("hello world")
}
