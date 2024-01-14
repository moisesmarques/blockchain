package main

import (
	"crypto/rsa"
	"fmt"
)

// Fak stores private and public keys
type Fak struct {
	Private rsa.PrivateKey
	// PrivateKey []uint8
	PublicKey []byte
	SubKeys   map[string]*FakSubKey
}

// NewFak creates and returns a FAK
func NewFak() *Fak {
	private, publicKey := newKeyPair()
	// private, privateKey, publicKey := newKeyPair()
	fak := Fak{*private, publicKey, make(map[string]*FakSubKey)}
	// fak := Fak{*private, privateKey, publicKey, make(map[string]*FakSubKey)}

	return &fak
}

func (ws Wallets) CreateFak(address string) string {

	wallet := *ws.Wallets[address]

	// get palette data
	palette_data := getPaletteData()
	// AP Key - get it and use it to create FAK
	// use shell creation process data
	shellc_data := createShell()

	newFak := NewFak()

	// initAccount(address)
	fakAddress := fmt.Sprintf("%s", newFak.GetAddress())
	fak_signature := spawnFAK(address, palette_data, wallet.Private)
	createDataForSharing(address, fak_signature, shellc_data, wallet.Private)
	// addFAK(address, fakAddress)
	// encryptedDataLength := getEncrpytedDataLength(address, fakAddress)

	// log.Printf("enc data length %d", encryptedDataLength)
	// Use the generated keys for FAK Creation and store it
	// Temporarily storing FAKs are disabled
	// wallet.Faks[fakAddress] = newFak

	return fakAddress
}

// GetAddress returns Fak address
func (fak Fak) GetAddress() []byte {
	pubKeyHash := HashPubKey(fak.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}
