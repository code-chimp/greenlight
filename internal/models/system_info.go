package models

type SystemInfo struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Revision    string `json:"revision,omitempty"`
	Status      string `json:"status"`
}
