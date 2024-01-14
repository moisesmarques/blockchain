package main

import (
	"crypto/rsa"
	"log"

	"github.com/google/uuid"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

// Output definitions
type ResourceKey struct {
	R_type         string         `json:"r_type"`
	RingPublicKey  kyber.Point    `json:"ring_public_key"`
	RingPrivateKey kyber.Scalar   `json:"ring_private_key"`
	RSAPublicKey   rsa.PublicKey  `json:"rsa_public_key"`
	RSAPrivateKey  rsa.PrivateKey `json:"rsa_private_key"`
}

type Shell struct {
	Id        string        `json:"id"`
	Resources []ResourceKey `json:"resources"`
	// Time          int        `json: "time"`
	// Total_expense int        `json: "total_expense"`
}

func getPaletteData() Palette {
	// Input Mocking
	paletteData := Palette{
		RequiredResources: []Resource{},
	}
	paletteData.RequiredResources = append(paletteData.RequiredResources, Resource{
		ActionID: "exec_hello",
		Type:     "cpu",
		Quantity: 1,
		Code: `
			// Sample data for running in the shell

			function sayHello() {
				console.log("Hello World!");
			}

			sayHello();
		`,
	})
	paletteData.RequiredResources = append(paletteData.RequiredResources, Resource{
		ActionID: "exec_hi",
		Type:     "gpu",
		Quantity: 2,
		Code: `
			// Sample data for running in the shell

			function sayHi() {
				console.log("Hi User!");
			}

			sayHi();
		`,
	})
	return paletteData
}

func createShell() *Shell {

	// Scan the input actions, based on the actions -> resources involved, create a role key for each new resource on the below format
	/*
		id: <shell_id>
		...
		AP_Key has to be added
		resources: [
			{
				type: cpu
				key: <key>
				...
			},
			{
				type: gpu
				key: <key>
			},
			{
				type: storage
				key: <key>
			},
			{
				type: admin
				key: <key>
			},
			{
				type: orchestration
				key: <key>
			}
		]
	*/

	// Input Mocking
	inputData := Palette{}

	inputData.RequiredResources = append(inputData.RequiredResources, Resource{
		Type:     "cpu",
		Quantity: 1,
	})
	inputData.RequiredResources = append(inputData.RequiredResources, Resource{
		Type:     "gpu",
		Quantity: 2,
	})

	// Output construction begins here
	shellData := new(Shell)
	id := uuid.New()
	shellData.Id = id.String()
	shellData.Resources = []ResourceKey{}
	shardTracking := make(map[string]bool)

	for _, s := range inputData.RequiredResources {
		suite := edwards25519.NewBlakeSHA256Ed25519()
		newPrivateKey := suite.Scalar().Pick(suite.RandomStream())
		newPublicKey := suite.Point().Mul(newPrivateKey, nil)

		for i := 0; i < s.Quantity; i++ {
			// Uses shardTracking map to add shard keys only once
			if !shardTracking[s.Type] {
				private, _ := newKeyPair()
				publicKey := private.PublicKey
				log.Printf("Keys generated for %s ", s.Type)
				shellData.Resources = append(shellData.Resources, ResourceKey{
					R_type:         s.Type,
					RingPublicKey:  newPublicKey,
					RingPrivateKey: newPrivateKey,
					RSAPublicKey:   publicKey,
					RSAPrivateKey:  *private,
				})
				shardTracking[s.Type] = true
			}
		}
	}
	// log.Printf("id =>> ", string(json.Marshal(shellData)))
	// privateKey := suite.Scalar().Pick(suite.RandomStream())
	return shellData
}
