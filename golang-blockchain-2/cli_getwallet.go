package main

import (
	"log"
)

func (cli *CLI) getWallet(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Wallet address is not valid")
	}

	wallets, _ := NewWallets(nodeID)

	wallet := wallets.Wallets[address]

	//fmt.Printf("Wallet '%+v'", wallet)

	PrintJson(wallet)
}
