package main

import (
	"github.com/aidarkhanov/nanoid"
)

type Palette struct {
	Id                  string
	Paints              map[string]*Paint
	TransactionHash     string
	Receiver            string
	RequiredResources   []Resource
	RequiredPermissions []string
	DecayFactors        map[string]*DecayFactorItem
	Comission           float32
}

type Resource struct {
	ActionID string `json:"action_id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
	Code     string `json:"code"`
}

type DecayFactorItem struct {
	Id    string
	Name  string
	Value string
}

// NewPalette creates and returns a Palette
func NewPalette() *Palette {
	palette := Palette{Id: nanoid.New()}
	palette.Paints = make(map[string]*Paint)
	palette.DecayFactors = make(map[string]*DecayFactorItem)
	palette.AddDecayFactor("expiry", "3600")

	return &palette
}

func NewDecayFactor(name string, value string) *DecayFactorItem {
	decayFactor := DecayFactorItem{
		Id:    nanoid.New(),
		Name:  name,
		Value: value,
	}

	return &decayFactor
}

func (palette *Palette) AddDecayFactor(name string, value string) string {

	decayFactor := NewDecayFactor(name, value)
	palette.DecayFactors[decayFactor.Id] = decayFactor

	return decayFactor.Id
}
