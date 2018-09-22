package models

// An Address entity that consits of City and State
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
