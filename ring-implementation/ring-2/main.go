/*

go mod init demo1
go mod tidy
go run main.go

*/

package main

import (
	"strconv"

	"fmt"
	"os"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/sign/anon"
)

func main() {

	m := "Hello"
	n := 3

	argCount := len(os.Args[1:])

	if argCount > 0 {
		m = string(os.Args[1])
	}
	if argCount > 1 {
		n, _ = strconv.Atoi(os.Args[2])
	}

	M := []byte(m)

	suite := edwards25519.NewBlakeSHA256Ed25519()

	// Create an anonymity set of random "public keys"
	X := make([]kyber.Point, n)
	for i := range X { // pick random points
		X[i] = suite.Point().Pick(suite.RandomStream())
	}
	fmt.Printf("network keys: %s\n\n", X)

	fmt.Println("==================================")

	// Make just one of them an actual public/private keypair (X[mine],x)
	mine := 1                                      // only the signer knows this
	x := suite.Scalar().Pick(suite.RandomStream()) // create a private key x
	X[mine] = suite.Point().Mul(x, nil)            // corresponding public key X

	// Generate the signature
	fmt.Printf("Private key of signer: %s\n\n", x)
	fmt.Printf("Public key of signer: %s\n\n", X[1])

	fmt.Println("==================================")

	fmt.Printf("network keys: %s\n\n", X)

	fmt.Println("==================================")

	sig := anon.Sign(suite, M, anon.Set(X), nil, mine, x)

	// fmt.Print("Signature - :\n" + hex.Dump(sig))
	fmt.Print("Signature:\n", sig)
	fmt.Println("")
	fmt.Println("==================================")

	x = suite.Scalar().Pick(suite.RandomStream())
	X[mine] = suite.Point().Mul(x, nil)

	fmt.Printf("network keys: %s\n\n", X)

	fmt.Println("==================================")

	// m = "Hellox"
	// M = []byte(m)

	// Verify the signature against the correct message
	tag, _ := anon.Verify(suite, M, anon.Set(X), nil, sig)

	if tag == nil || len(tag) != 0 {
		panic("Verify returned wrong tag")
	}
	fmt.Printf("\n\nSignature has been verified")
	fmt.Println("")

}
