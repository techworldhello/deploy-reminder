package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/techworldhello/deploy-reminder/webhook-reciever/pkg/store"
	"github.com/techworldhello/deploy-reminder/webhook-reciever/pkg/webhook_parser"
	"log"
)

type Webhook struct {
	Body string `json:"body"`
}

func WebhookReciever(ctx context.Context, event Webhook) (string, error) {
	// check if header user-agent is Buildkite-Request
	w := webhook_parser.NewBuildKite()
	if err := w.Parse(event.Body); err != nil {
		log.Fatalf("error parsing data: %+v", err)
	}
	s := store.New()
	if err := s.SaveData(w.JobFinished); err != nil {
		log.Fatalf("error saving data: %+v", err)
	}
	return fmt.Sprintf("Webhook event: %s", event), nil
}

func main() {
	lambda.Start(WebhookReciever)
}