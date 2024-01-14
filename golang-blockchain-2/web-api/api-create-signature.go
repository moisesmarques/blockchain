package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
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

func main() {

	router := gin.Default()
	router.GET("/shards", getShards)
	router.GET("/shards/:id", getShardByID)
	router.POST("/shards", postShells)
	router.POST("/createSignature", createSig)
	router.POST("/verifySignature", verifySig)
	router.Run("localhost:8080")
}

func getShards(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, shards)
}

func postShells(c *gin.Context) {
	var newShell shard

	if err := c.BindJSON(&newShell); err != nil {
		return
	}

	// Add the new shard to the slice.
	newShell.RingSignaturePublicKey = suite.Point().Pick(suite.RandomStream())
	newShell.RingSignaturePrivateKey = suite.Scalar().Pick(suite.RandomStream())
	shards = append(shards, newShell)
	c.IndentedJSON(http.StatusCreated, newShell)
}

func getShardByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for _, a := range shards {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "shard not found"})
}

func createSig(c *gin.Context) {

	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	jsonBody := string(bodyAsByteArray)
	fmt.Println("jsonBody: ", jsonBody)

	id := gjson.Get(jsonBody, "id").String()

	fmt.Println("id: ", id)
	keys := make([]kyber.Point, 0)
	fmt.Println("keys initial: ", keys)

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for index, a := range shards {
		if a.ID == id {
			//c.IndentedJSON(http.StatusOK, a)
			//return
			privateKeyIndex = index
			privateKey = a.RingSignaturePrivateKey
			keys = append(keys, suite.Point().Mul(privateKey, nil))

		} else {
			keys = append(keys, a.RingSignaturePublicKey)
		}

	}
	fmt.Println("create signature keys: ", keys)

	fmt.Println("privateKey: ", privateKey)
	fmt.Println("privateKeyIndex: ", privateKeyIndex)

	cignature := createSignature(keys, privateKey, privateKeyIndex, message)
	signature := string(cignature[:])
	c.IndentedJSON(http.StatusCreated, signature)

}
func verifySig(c *gin.Context) {

	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	jsonBody := string(bodyAsByteArray)
	fmt.Println("jsonBody: ", jsonBody)

	signatureString := gjson.Get(jsonBody, "signature").String()
	signature := []byte(signatureString)
	fmt.Println("signature: ", signature)

	keys := make([]kyber.Point, 0)
	fmt.Println("keys initial: ", keys)

	// Loop through the list of shards, looking for
	// an shard whose ID value matches the parameter.
	for index, a := range shards {
		if index == privateKeyIndex {
			privateKey = a.RingSignaturePrivateKey
			keys = append(keys, suite.Point().Mul(privateKey, nil))

		} else {
			keys = append(keys, a.RingSignaturePublicKey)
		}

	}
	fmt.Println("verifySig keys: ", keys)
	verifySignature(keys, ringSignature, message)
	// verifySignature(keys, signature, message)

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

func verifySignature(keys []kyber.Point, circularSig []byte, message string) {

	// Modifying the given slice

	M := []byte(message)
	suite := edwards25519.NewBlakeSHA256Ed25519()

	verificationTag, _ := anon.Verify(suite, M, anon.Set(keys), nil, circularSig)
	if verificationTag == nil || len(verificationTag) != 0 {
		panic("Verify returned wrong tag2")
	}
	fmt.Println("\n\n verifySignature circularSig has been verified", verificationTag)

}
