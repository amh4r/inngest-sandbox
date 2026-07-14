package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/inngest/inngestgo"
	"github.com/inngest/inngestgo/step"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("error loading .env: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "dev"
	}

	client, err := inngestgo.NewClient(inngestgo.ClientOpts{
		AppID:      "sandbox-go-connect",
		AppVersion: inngestgo.Ptr(appVersion),
		Logger:     slog.Default(),
	})
	if err != nil {
		log.Fatalf("error creating Inngest client: %v", err)
	}

	_, err = inngestgo.CreateFunction(
		client,
		inngestgo.FunctionOpts{
			ID:   "fn-1",
			Name: "Go Connect example",
		},
		inngestgo.EventTrigger("event-1", nil),
		func(ctx context.Context, input inngestgo.Input[map[string]any]) (any, error) {
			message, err := step.Run(ctx, "build-message", func(ctx context.Context) (string, error) {
				return fmt.Sprintf("Hello from Go Connect. Event data: %v", input.Event.Data), nil
			})
			if err != nil {
				return nil, err
			}

			return map[string]any{
				"message": message,
			}, nil
		},
	)
	if err != nil {
		log.Fatalf("error creating function: %v", err)
	}

	maxWorkerConcurrency := int64(10)
	connection, err := inngestgo.Connect(ctx, inngestgo.ConnectOpts{
		Apps:                 []inngestgo.Client{client},
		InstanceID:           inngestgo.Ptr(workerInstanceID()),
		MaxWorkerConcurrency: &maxWorkerConcurrency,
	})
	if err != nil {
		log.Fatalf("error connecting worker: %v", err)
	}

	log.Printf("go-connect worker connected: state=%s", connection.State())
	<-ctx.Done()

	log.Println("go-connect worker shutting down")
	if err := connection.Close(); err != nil {
		log.Fatalf("error closing worker connection: %v", err)
	}
	log.Printf("go-connect worker closed: state=%s", connection.State())
}

func workerInstanceID() string {
	if id := os.Getenv("INNGEST_CONNECT_INSTANCE_ID"); id != "" {
		return id
	}

	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		return "go-connect-worker"
	}

	return hostname
}
