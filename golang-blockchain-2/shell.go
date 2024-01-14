package main

import (
	"time"
)

//"github.com/golang/protobuf/ptypes/timestamp"

type Shell2 struct {
	TransactionHash   string
	ActionNames       []string
	ValidatorTypes    []string
	TotalExpense      float32
	Receiver          string
	Sender            string
	Permissions       []string
	ConsumedResources []string
	Time              time.Time
	Nonce             int
	Duration          int
	Results           []*ShellResult
}
