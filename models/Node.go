package models

type Node struct {
	ID         string  `json:"_id"`
	VRAM       float64 `json:"vram"`
	CudaDevice string  `json:"cuda_device"`
	GPUCount   int     `json:"gpus"`
	Type       string  `json:"mode"` // either image or text
	Model      string  `json:"model,omitempty"`
	MachineID  string  `json:"machineId,omitempty"`
	OwnerID    string  `json:"ownerId,omitempty"`
}
