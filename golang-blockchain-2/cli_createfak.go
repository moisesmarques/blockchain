package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createFak(wallet, nodeID string) {
	if !ValidateAddress(wallet) {
		log.Panic("ERROR: Wallet address is not valid")
	}

	wallets, _ := NewWallets(nodeID)
	fakAddress := wallets.CreateFak(wallet)
	wallets.SaveToFile(nodeID)

	privateKey := suite.Scalar().Pick(suite.RandomStream())

	ringSignature := createSigUsingFAK(fakAddress, privateKey)
	//isVerified := verifySigUsingFAK(fakAddress, privateKey)

	fmt.Printf("RingSignature generated '%v' \n", ringSignature[0])

	isVerified := verifySigUsingFAKAndSignature(fakAddress, privateKey, ringSignature)
	fmt.Printf("FAK generated '%s'\n", fakAddress)
	fmt.Printf("RingSignature generated '%v' \n", ringSignature)
	fmt.Printf("RingSignature isVerified '%v' \n", isVerified)

}
