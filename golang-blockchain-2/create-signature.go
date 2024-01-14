// /*

// go mod init demo1
// go mod tidy
// go run main.go

// */

package main

import (
	"fmt"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/sign/anon"
)

func createSignature2(keys []kyber.Point, privateKey kyber.Scalar, functionalAccessKey string) []byte {

	// Modifying the given slice

	M := []byte(functionalAccessKey)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	circularSig := anon.Sign(suite, M, anon.Set(keys), nil, len(keys)-1, privateKey)
	fmt.Println("createSignature2 circularSig: ", circularSig)

	fmt.Println("signatures: ", keys)
	fmt.Println("privateKey: ", privateKey)
	return circularSig

}

func verifySignature2(keys []kyber.Point, circularSig []byte, functionalAccessKey string) {

	// Modifying the given slice

	M := []byte(functionalAccessKey)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	verificationTag, _ := anon.Verify(suite, M, anon.Set(keys), nil, circularSig)
	if verificationTag == nil || len(verificationTag) != 0 {
		panic("Verify returned wrong tag2")
	}
	fmt.Println("\n\n verifySignature circularSig has been verified", verificationTag)

}

func test() {

	functionalAccessKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImNoaXJhZyBwdXJvaGl0IiwiaWF0IjoxNTE2MjM5MDIyfQ.7o_alUu9fG0d8MyZ4iQnDJIxy9ENBG-wgpNchhj2gIM"
	suite := edwards25519.NewBlakeSHA256Ed25519()

	networkKey := suite.Point().Pick(suite.RandomStream())
	roleKey := suite.Point().Pick(suite.RandomStream())
	privateKey := suite.Scalar().Pick(suite.RandomStream())

	keys := make([]kyber.Point, 3)
	keys[0] = networkKey
	keys[1] = roleKey
	keys[2] = suite.Point().Mul(privateKey, nil)

	circularSig2 := createSignature2(keys, privateKey, functionalAccessKey)
	fmt.Print("circularSig2 :\n", circularSig2)

	verifySignature2(keys, circularSig2, functionalAccessKey)

	fmt.Println("==================================")
	/*
		M := []byte(functionalAccessKey)
		user := 2
		circularSig := anon.Sign(suite, M, anon.Set(keys), nil, user, privateKey)

		fmt.Print("Signature :\n", circularSig)
		fmt.Println("==================================")

		functionalAccessKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImNoaXJhZyBwdXJvaGl0IiwiaWF0IjoxNTE2MjM5MDIyfQ.7o_alUu9fG0d8MyZ4iQnDJIxy9ENBG-wgpNchhj2gIM"
		M = []byte(functionalAccessKey)

		verificationTag, _ := anon.Verify(suite, M, anon.Set(keys), nil, circularSig)
		if verificationTag == nil || len(verificationTag) != 0 {
			panic("Verify returned wrong tag2")
		}
		fmt.Println("\n\n circularSig has been verified", verificationTag)
	*/
}
