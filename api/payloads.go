package api

// LoginPayload /authenticate
type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TokenAuthentication result
type TokenAuthentication struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
	APIKey  string `json:"apiKey"`
}

type FiberValidErr struct {
	Success bool             `json:"success"`
	Errors  []*ErrorResponse `json:"errors"`
}
type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag,omitempty"`
	Value       string `json:"value"`
}
