package main

import (
	"fmt"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/sign/anon"
)

var suite edwards25519.SuiteEd25519
var privateKey kyber.Scalar
var privateKeyIndex int
var message = ""
var ringSignature []byte

// shard represents data about a record shard.
type shard struct {
	ID                      string       `json:"id"`
	Title                   string       `json:"title"`
	RingSignaturePublicKey  kyber.Point  `json:"ringSignaturePrivateKey"`
	RingSignaturePrivateKey kyber.Scalar `json:"ringSignaturePublicKey"`
}

// shards slice to seed record shard data.
var shards = []shard{

	{ID: "1", Title: "GPU",
		RingSignaturePublicKey:  suite.Point().Pick(suite.RandomStream()),
		RingSignaturePrivateKey: suite.Scalar().Pick(suite.RandomStream()),
	},
	{ID: "2", Title: "CPU",
		RingSignaturePublicKey:  suite.Point().Pick(suite.RandomStream()),
		RingSignaturePrivateKey: suite.Scalar().Pick(suite.RandomStream()),
	},
	{ID: "3", Title: "STORAGE",
		RingSignaturePublicKey:  suite.Point().Pick(suite.RandomStream()),
		RingSignaturePrivateKey: suite.Scalar().Pick(suite.RandomStream()),
	},
}

func createSigUsingFAK(fak string, privateKey kyber.Scalar) []byte {

	id := "2"

	fmt.Println("id: ", id)
	keys := make([]kyber.Point, 0)
	fmt.Println("keys initial: ", keys)

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for _, a := range shards {
		keys = append(keys, a.RingSignaturePublicKey)
	}
	keys = append(keys, suite.Point().Mul(privateKey, nil))

	fmt.Println("create signature keys: ", keys)
	fmt.Println("privateKey: ", privateKey)

	cignature := createSignature(keys, privateKey, (len(keys) - 1), fak)
	return cignature

}

func verifySigUsingFAK(fak string, privateKey kyber.Scalar) bool {

	keys := make([]kyber.Point, 0)
	fmt.Println("keys initial: ", keys)

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for _, a := range shards {
		keys = append(keys, a.RingSignaturePublicKey)
	}
	fmt.Println("verifySig keys before private key: ", keys)
	keys = append(keys, suite.Point().Mul(privateKey, nil))
	fmt.Println("verifySig keys after private key: ", keys)
	isVerified := verifySignature(keys, ringSignature, fak)
	return isVerified
}

func verifySigUsingFAKAndSignature(fak string, privateKey kyber.Scalar, signature []byte) bool {

	keys := make([]kyber.Point, 0)
	fmt.Println("keys initial: ", keys)

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for _, a := range shards {
		keys = append(keys, a.RingSignaturePublicKey)
	}
	fmt.Println("verifySig keys before private key: ", keys)
	keys = append(keys, suite.Point().Mul(privateKey, nil))
	fmt.Println("verifySig keys after private key: ", keys)
	isVerified := verifySignature(keys, signature, fak)
	return isVerified
}

func createSignature(keys []kyber.Point, privateKey kyber.Scalar, privateKeyIndex int, message string) []byte {

	// Modifying the given slice

	M := []byte(message)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	ringSignature = anon.Sign(suite, M, anon.Set(keys), nil, privateKeyIndex, privateKey)

	fmt.Println("createSignature2 circularSig: ", ringSignature)
	fmt.Println("circularSig: ", ringSignature)

	return ringSignature

}

func verifySignature(keys []kyber.Point, circularSig []byte, message string) bool {

	// Modifying the given slice

	M := []byte(message)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	verificationTag, _ := anon.Verify(suite, M, anon.Set(keys), nil, circularSig)
	if verificationTag == nil || len(verificationTag) != 0 {
		//panic("Verify returned wrong tag2")
		return false
	}
	fmt.Println("\n\n verifySignature circularSig has been verified", verificationTag)
	return true

}
