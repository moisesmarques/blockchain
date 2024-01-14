package main

import (
	// "crypto/rsa"
	"fmt"
	"log"

	// "math/big"
	"math/rand"
	"strings"
	"time"

	// "github.com/rogercoll/rsignatures"
	"go.dedis.ch/kyber/v3"
)

func (cli *CLI) createAPKeySignature(wallets string,
	message string,
	nodeID string) {

	wallets_ := strings.Split(wallets, ",")

	partyKeys := make([]kyber.Point, 0)

	rand.Seed(time.Now().UnixNano())

	signerRound := rand.Intn(len(wallets_))

	fmt.Printf("signerRound: %d\n", signerRound)

	var signerKey kyber.Scalar

	for i, wallet := range wallets_ {
		fmt.Printf("wallet %s %d\n", wallet, i)
		walletObj := getWallet(wallet, nodeID)
		fmt.Printf("RingPublicKey %d\n", walletObj.RingPublicKey)

		partyKeys = append(partyKeys, walletObj.RingPublicKey)

		// if i == signerRound {
		// 	signerKey = *walletObj.RingPrivateKey
		// }
	}

	message_ := "test"

	sig := createSignature(partyKeys, signerKey, signerRound, message_)

	//rsaRing := rsignatures.NewRSARing(partyKeys, signerKey)

	res := verifySignature(partyKeys, sig, message_)

	//seed, sig, err := rsaRing.Sign([]byte(message_), signerRound)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("The signature was generated: %s\n", sig)

	fmt.Printf("Signature validated: %t\n", res)

}

func (cli *CLI) verifyAPKeySignature(wallets string,
	signature string,
	//seed string,
	message string,
	nodeID string) {

	wallets_ := strings.Split(wallets, ",")
	signature_ := []byte(signature)

	// seed_ := new(big.Int)
	// seed_, ok := seed_.SetString(seed, 10)
	// if !ok {
	// 	fmt.Println("Error converting seed to big int.")
	// 	return
	// }

	// fmt.Printf("Verify AP Key seed %s\n", seed_)

	partyKeys := make([]kyber.Point, 0)

	//signerRound := rand.Intn(len(wallets_))

	// var signerKey *rsa.PrivateKey

	for _, wallet := range wallets_ {

		walletObj := getWallet(wallet, nodeID)
		partyKeys = append(partyKeys, walletObj.RingPublicKey)

		// if i == signerRound {
		// 	signerKey = &walletObj.Private
		// }
	}

	// rsaRing := rsignatures.NewRSARing(partyKeys, signerKey)

	res := verifySignature(partyKeys, signature_, message)

	fmt.Printf("Signature validated: %t\n", res)

}

func getWallet(address string, nodeID string) Wallet {

	if !ValidateAddress(address) {
		log.Panic("ERROR: Wallet address is not valid")
	}

	wallets, _ := NewWallets(nodeID)

	wallet := wallets.Wallets[address]

	return *wallet
}
