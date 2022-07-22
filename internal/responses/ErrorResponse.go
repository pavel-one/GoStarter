package responses

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
