package response

type ErrorMessage struct {
	Message  string `json:"message"`
	Required string `json:"required,omitempty"`
	Cause    string `json:"cause,omitempty"`
}
