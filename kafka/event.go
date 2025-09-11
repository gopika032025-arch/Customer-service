package kafka

type CustomerEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Type  string `json:"type"` // "created", "updated", "deactivated"
}

const (
	CustomerCreatedTopic     = "customer-created"
	CustomerUpdatedTopic     = "customer-updated"
	CustomerDeactivatedTopic = "customer-deactivated"
)
