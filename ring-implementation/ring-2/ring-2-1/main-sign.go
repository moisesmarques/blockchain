/*

go mod init demo1
go mod tidy
go run main.go

*/

package main

import (
	"fmt"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/sign/anon"
)

func main() {

	functionalAccessKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImNoaXJhZyBwdXJvaGl0IiwiaWF0IjoxNTE2MjM5MDIyfQ.7o_alUu9fG0d8MyZ4iQnDJIxy9ENBG-wgpNchhj2gIM"
	M := []byte(functionalAccessKey)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	networkKey := suite.Point().Pick(suite.RandomStream())
	roleKey := suite.Point().Pick(suite.RandomStream())
	userKey := suite.Scalar().Pick(suite.RandomStream())

	user := 2
	keys := make([]kyber.Point, 3)

	keys[0] = networkKey
	keys[1] = roleKey
	keys[2] = suite.Point().Mul(userKey, nil)

	circularSig := anon.Sign(suite, M, anon.Set(keys), nil, user, userKey)

	fmt.Print("Signature :\n", circularSig)
	fmt.Println("==================================")

	functionalAccessKey = "eyJhbciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImNoaXJhZyBwdXJvaGl0IiwiaWF0IjoxNTE2MjM5MDIyfQ.7o_alUu9fG0d8MyZ4iQnDJIxy9ENBG-wgpNchhj2gIM"
	M = []byte(functionalAccessKey)

	verificationTag, _ := anon.Verify(suite, M, anon.Set(keys), nil, circularSig)
	if verificationTag == nil || len(verificationTag) != 0 {
		panic("Verify returned wrong tag2")
	}
	fmt.Println("\n\n circularSig has been verified", verificationTag)

}
