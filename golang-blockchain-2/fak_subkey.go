package main

import (
	"fmt"

	"github.com/aidarkhanov/nanoid"
)

// Fak stores private and public keys
type FakSubKey struct {
	Id string
}

// NewFak creates and returns a FAK SubKey
func NewFakSubKey() *FakSubKey {
	subKeyAddress := fmt.Sprintf("%s", nanoid.New())
	fakSubKey := FakSubKey{subKeyAddress}

	return &fakSubKey
}

func (fak Fak) CreateFakSubKey() *FakSubKey {

	newFakSubkey := NewFakSubKey()

	fak.SubKeys[newFakSubkey.Id] = newFakSubkey

	return newFakSubkey
}
