package webhook_parser

import "time"

type JobFinished struct {
	Event string `json:"event"`
	Id string `json:"id,omitempty"`
	Job   struct {
		CreatedAt          time.Time   `json:"created_at"`
		ScheduledAt        time.Time   `json:"scheduled_at"`
		RunnableAt         time.Time   `json:"runnable_at"`
		StartedAt          time.Time   `json:"started_at"`
		FinishedAt         time.Time   `json:"finished_at"`
		Retried            bool        `json:"retried"`
		RetriedInJobID     interface{} `json:"retried_in_job_id"`
		RetriesCount       interface{} `json:"retries_count"`
	} `json:"job"`
	Build struct {
		ID           string      `json:"id"`
		Number       int         `json:"number"`
		State        string      `json:"state"`
		Blocked      bool        `json:"blocked"`
		BlockedState string      `json:"blocked_state"`
		Message      string      `json:"message"`
		Commit       string      `json:"commit"`
		Branch       string      `json:"branch"`
		Tag          interface{} `json:"tag"`
		Source       string      `json:"source"`
		Author       interface{} `json:"author"`
		Creator      struct {
			ID        string    `json:"id"`
			Name      string    `json:"name"`
			Email     string    `json:"email"`
			AvatarURL string    `json:"avatar_url"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"creator"`
		CreatedAt   time.Time   `json:"created_at"`
		ScheduledAt time.Time   `json:"scheduled_at"`
		StartedAt   time.Time   `json:"started_at"`
		FinishedAt  interface{} `json:"finished_at"`
	} `json:"build"`
	Sender struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"sender"`
}
