package shell

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)

// Create a Shell
// params:
//		PalletJson : (string)  Requirements Pallet in Json format
// output:
//		(string) resulting Shell information in JSON format
//

func CreateShell(PalletJson string) {
	//if !ValidateAddress(wallet) {
	//	log.Panic("ERROR: Wallet address is not valid")
	//}

	// 1. Define a struct for the Pallet
	// 2. Parse the Pallet json string into the Struct
	// 3. Process the Pallet struct to find resource requirements
	// 4. Allocate resources from the Global Resourse Pool(Maintain in Arrays, fetch using a function) to the Shell (we will allocate all to begin with)
	// 5. Create a FAK based on allocated Resources (Need to collaborate with FAK team for this)
	// 6. Hand over the Shell Schema and FAK to Admin/s of the Shell. (may come next week)
	globalResourcePool := GetResourcePool()

	//go run . createshell -pallet "{\"user_id\":\"abc\", \"Function_name\":\"RenderScene\", \"Storage_requirements\": {\"type\":\"ssd\", \"count\":2, , \"space\":1000000000, , \"hydration\":\"no\"} }"

	var inputPallet Palette
	json.Unmarshal([]byte(PalletJson), &inputPallet)

	//pallet.VerifyPallet(PalletJson)
	fmt.Printf("\n\n------------------------------------------------------\n")
	fmt.Printf("\nReceived Pallet from User ID: '%s'", inputPallet.User_id)
	fmt.Printf("\nRequire:\n\t'%d' CPU with %d bytes of memory ", inputPallet.Cpu_requirements.Count, inputPallet.Cpu_requirements.Memory)
	fmt.Printf("\n\t'%d' GPU with min '%dMhz' frequency and '%s' feature ", inputPallet.Gpu_Requirements.Count, inputPallet.Gpu_Requirements.MinFrequency, inputPallet.Gpu_Requirements.Features)
	fmt.Printf("\n\t'%d' Storage nodes with %d bytes of '%s' space", inputPallet.Storage_requirements.Count, inputPallet.Storage_requirements.Space, inputPallet.Storage_requirements.Type)

	fmt.Printf("\n\nCreating Shell....")
	fmt.Printf("\n\nAllocating resources....")
	newShell := new(Shell)
	newShell.Palette = inputPallet
	newShell.Hash = createShellHash()
	newShell.Computational_resources = globalResourcePool
	fmt.Printf("[DONE]")
	fmt.Printf("\n\n------------------------------------------------------\n")
	out, err := json.Marshal(newShell)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(out))
	fmt.Printf("\n\n")
}

func createShellHash() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
