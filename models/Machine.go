package models

type Machine struct {
	Nodes     []Node `json:"nodes"`
	MachineID string `json:"machineId"`
	OwnerID   string `json:"ownerId"`
}
