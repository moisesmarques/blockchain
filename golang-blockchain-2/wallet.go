package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"log"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	Private             rsa.PrivateKey
	PublicKey           []byte
	RingPublicKey       kyber.Point
	RingPrivateKey      kyber.Scalar
	Faks                map[string]*Fak
	RingPublicKeyBytes  []byte
	RingPrivateKeyBytes []byte
}

func (w *Wallet) SetRingPublicKey() {

	fmt.Printf("SetRingPublicKey %d\n", w.RingPublicKey)

	a_by, err := w.RingPublicKey.Data()
	if err != nil {
		log.Fatal(err)
	}
	w.RingPublicKeyBytes = a_by
}

func (w *Wallet) SetRingPrivateKey() {
	//w.RingPrivateKeyBytes = w.RingPrivateKey
	fmt.Printf("SetRingPrivateKey %d\n", w.RingPrivateKey)
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {

	// private, privateKey, publicKey := newKeyPair()
	private, publicKey := newKeyPair()
	log.Printf("RSA Keys generated")

	privateRng, publicRng := newRingKeyPair()
	log.Printf("Kyber Keys generated private %d public %d", privateRng, publicRng)
	// wallet := Wallet{*private, privateKey, publicKey, make(map[string]*Fak)}
	wallet := Wallet{Private: *private,
		PublicKey:      publicKey,
		RingPublicKey:  *publicRng,
		RingPrivateKey: *privateRng,
		Faks:           make(map[string]*Fak)}

	wallet.SetRingPublicKey()
	wallet.SetRingPrivateKey()

	log.Printf("Wallet generated")
	return &wallet
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

// func newKeyPair() (*rsa.PrivateKey, []byte, []byte) {
func newKeyPair() (*rsa.PrivateKey, []byte) {
	log.Printf("generating key..")
	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("generating key complete")
	publickey := &private.PublicKey
	// privateKey := x509.MarshalPKCS1PrivateKey(private)
	pubKey, pubKeyErr := x509.MarshalPKIXPublicKey(publickey)
	if pubKeyErr != nil {
		log.Panic(pubKeyErr)
	}
	// return private, privateKey, pubKey
	return private, pubKey
}

func newRingKeyPair() (*kyber.Scalar, *kyber.Point) {
	log.Printf("generating ring key pair..")
	suite := edwards25519.NewBlakeSHA256Ed25519()
	newPublicKey := suite.Point().Pick(suite.RandomStream())
	newPrivateKey := suite.Scalar().Pick(suite.RandomStream())
	// return privateKey, pubKey
	return &newPrivateKey, &newPublicKey
}
