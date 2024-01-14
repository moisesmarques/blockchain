package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet_%s.dat"

// Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallet
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile(nodeID)

	return &wallets, err
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}

// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetWallet returns a Wallet by its address
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) CreateFakSubKey(walletAddress string, fakAddress string) string {

	wallet := ws.Wallets[walletAddress]

	fak := wallet.Faks[fakAddress]

	fakSubKey := fak.CreateFakSubKey()

	return fakSubKey.Id
}

// LoadFromFile loads wallets from the file
func (ws *Wallets) LoadFromFile(nodeID string) error {

	walletFile := fmt.Sprintf(walletFile, nodeID)

	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	// gob.Register(elliptic.P256())
	// decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	// err = decoder.Decode(&wallets)
	// if err != nil {
	// 	log.Panic(err)
	// }

	_ = json.Unmarshal([]byte(fileContent), &wallets)

	ws.Wallets = wallets.Wallets

	return nil
}

// SaveToFile saves wallets to a file
func (ws Wallets) SaveToFile(nodeID string) {

	fmt.Print("salvando wallets")

	walletFile := fmt.Sprintf(walletFile, nodeID)

	// var content bytes.Buffer
	// gob.Register(elliptic.P256())

	// encoder := gob.NewEncoder(&content)
	// err := encoder.Encode(ws)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	// if err != nil {
	// 	log.Panic(err)
	// }

	file, _ := json.MarshalIndent(ws, "", " ")
	ioutil.WriteFile(walletFile, file, 0644)

}
