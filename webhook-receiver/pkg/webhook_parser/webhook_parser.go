package webhook_parser

import (
	"encoding/json"
	"log"
)

type BuildKite struct {
	JobFinished
}

func NewBuildKite() *BuildKite {
	return &BuildKite{JobFinished{}}
}

func (b *BuildKite) Parse(webhookEvent string) error {
	// check if header x-buildkite-event is job.finished
	log.Printf("In parse")
	if err := json.Unmarshal([]byte(webhookEvent), &b.JobFinished); err != nil {
		return err
	}
	return nil
}