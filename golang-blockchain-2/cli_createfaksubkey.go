package main

import (
	"fmt"
	"log"
)

func (cli *CLI) CreateFakSubKey(wallet, fak, nodeID string) {
	if !ValidateAddress(wallet) {
		log.Panic("ERROR: Wallet address is not valid")
	}

	wallets, _ := NewWallets(nodeID)

	fakSubKeyId := wallets.CreateFakSubKey(wallet, fak)

	wallets.SaveToFile(nodeID)

	fmt.Printf("FAK subKey generated '%s'", fakSubKeyId)
}
