package zendesk

// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/ticket_audits#the-via-object
type Via struct {
	Channel	*string	`json:"channel,omitempty"`
        Source	*Source	`json:"source,omitempty"`
}
type Source struct{
	To	*To	`json:"to,omitempty"`
	From	*From	`json:"from,omitempty"`
}
type To	struct{
	Name	*string	`json:"name,omitempty"`
	Address	*string	`json:"address,omitempty"`
}
type From struct{
	Address	*string	`json:"address,omitempty"`
	Original  []string	`json:"original_recipients,omitempty"`
}
