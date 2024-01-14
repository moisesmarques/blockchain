package main

import (
	"time"

	"github.com/aidarkhanov/nanoid"
)

type Paint struct {
	Id                   string
	ActionName           string
	Permissions          []string
	ResourceRequirements []ResourceRequirement
	AppRequirements      []AppRequirement
	NetworkRequirements  []NetworkRequirement
	Time                 time.Time
	ValidatorType        string
	Expense              string
	Sender               string
}

type ResourceRequirement struct {
	ActionID string `json:"action_id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
}
type NetworkRequirement struct {
	Complexity      int
	ValidatorType   string
	StorageCapacity int
	Hydration       int
}

type AppRequirement struct {
	Comission          string
	ValidatorType      string
	PaymentDesignation string
	DecayRate          int
}

// NewPaint creates and returns a Paint
func NewPaint() *Paint {
	paint := Paint{Id: nanoid.New()}
	return &paint
}
