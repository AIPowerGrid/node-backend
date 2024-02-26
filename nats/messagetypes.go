package nats

type modelResponse struct {
	Model string `json:"model"`
	Type  string `json:"type"`
}

type ComfyPayload struct {
	Prompt string `json:"prompt"`
	Seed   string `json:"seed"`
}
type JSResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
}
