package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/random"
)

var (
	fakHolder map[string]map[string]*FAKContent
)

type FAKContent struct {
	Data          *Palette        `json:"data"` // Palette data
	DataHash      [32]byte        `json:"data_hash"`
	Encryptions   []EncryptedCore `json:"encryptions"`
	ShellResults  *Shell          `json:"shell_results"`
	RingSignature string          `json:"ring_signature"`
}

type EncryptedCore struct {
	Id               string `json:"id"`
	EncAES           []byte `json:"enc_AES"`
	EncData          []byte `json:"enc_data"`
	EncDataSignature []byte `json:"enc_data_signature"`
}

func initAccount(account_id string) {
	fakHolder = make(map[string]map[string]*FAKContent)
	fakHolder[account_id] = make(map[string]*FAKContent)
}

func addFAK(account_id string, fak string) {
	fakHolder[account_id][fak] = new(FAKContent)
	log.Printf("FAK added to the account - %s", account_id)
}

func addData(account_id string, fak string, data *Palette, data_hash [32]byte) {
	fakHolder[account_id][fak].Data = data
	fakHolder[account_id][fak].DataHash = data_hash
	log.Printf("Palette data added to the account FAK")
}

func addShellResults(account_id string, fak string, data *Shell) {
	fakHolder[account_id][fak].ShellResults = data
	log.Printf("Shell data added to the account FAK")
}

func addEncData(account_id string, fak string, data EncryptedCore) {
	fakHolder[account_id][fak].Encryptions = append(fakHolder[account_id][fak].Encryptions, data)
	log.Printf("Encrypted data added to the account FAK")
}

func spawnFAK(account_id string, palette_data Palette, private_key rsa.PrivateKey) string {
	/*
		Parse the palette_data to bytes
		Create hash with sha256
		Sign and return the data => FAK
	*/
	initAccount(account_id)
	log.Printf("account initialized for FAK generation")
	palette_obj, errorOnMarshal := json.Marshal(palette_data)
	if errorOnMarshal != nil {
		log.Printf("got errorOnMarshal: ", errorOnMarshal)
	}
	palette_obj_stringified := string(palette_obj)
	log.Printf("chk ==> ", palette_obj_stringified)
	hashed := sha256.Sum256([]byte(palette_obj_stringified))
	signature, err := rsa.SignPKCS1v15(rand.Reader, &private_key, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	signature_str := string(signature)
	log.Printf("signature created")
	addFAK(account_id, signature_str)
	addData(account_id, signature_str, &palette_data, hashed)
	return signature_str
}

func createDataForSharing(account_id string, fak string, shell_data *Shell, private_key rsa.PrivateKey) {
	accountPublicKey := private_key.PublicKey
	paletteData := fakHolder[account_id][fak].Data
	addShellResults(account_id, fak, shell_data)
	for _, s := range paletteData.RequiredResources {
		resource_values := func(resource_obj ResourceKey) bool { return resource_obj.R_type == s.Type }
		filtered_value := filterShellData(shell_data.Resources, resource_values)
		s_obj, sMarshal := json.Marshal(s)
		if sMarshal != nil {
			log.Printf("got errorOnMarshal: ", sMarshal)
		}
		log.Printf("check this -> ", string(s_obj))
		// Encryption starts here
		if len(filtered_value) > 0 {
			role_public_key := filtered_value[0].RSAPublicKey
			//role_kyber_public_key := filtered_value[0].RingPublicKey
			//role_kyber_pvt_key := filtered_value[0].RingPrivateKey
			// Do encrpytion with account key & AES
			key := randBytes(256 / 8)
			block, err := aes.NewCipher(key)
			if err != nil {
				fmt.Printf("Error reading key: %s\n", err.Error())
				os.Exit(1)
			}

			fmt.Printf("Key: %s\n", hex.EncodeToString(key))

			gcm, err := cipher.NewGCM(block)
			if err != nil {
				fmt.Printf("Error initializing AEAD: %s\n", err.Error())
				os.Exit(1)
			}

			nonceSize := gcm.NonceSize()
			nonce := randBytes(nonceSize)
			c := gcm.Seal(nil, nonce, []byte(s.Code), nil)
			enc_data := append(nonce, c...)
			// log.Printf("encrypted value: %s", string(enc_data))
			log.Printf("encrypted the value")
			PrintJson(enc_data)
			// Encrypt AES with account key
			encryptedAES, encrpytionAESErr := rsa.EncryptPKCS1v15(rand.Reader, &accountPublicKey, key)
			if encrpytionAESErr != nil {
				fmt.Printf("Error encrypting AES key: %s\n", encrpytionAESErr.Error())
				os.Exit(1)
			}
			log.Printf("encrypted the AES with account key")
			// log.Printf("Enc AES %s", string(encryptedAES))
			PrintJson(encryptedAES)
			// Grant permission to the related key
			encryptedRoleAES, encrpytionRoleAESErr := rsa.EncryptPKCS1v15(rand.Reader, &role_public_key, key)
			// encryptedRoleAES, encrpytionRoleAESErr := ElGamalEncrypt(edwards25519, role_kyber_public_key, key)
			if encrpytionRoleAESErr != nil {
				fmt.Printf("Error encrypting AES key: %s\n", encrpytionRoleAESErr.Error())
				os.Exit(1)
			}
			log.Printf("encryptedRoleAES -> ", encryptedRoleAES)
			// Send AES key to the related nodes to use it for decrypting the encrypted data
			// Generate Enc Data Signature with account's private key
			hashed := sha256.Sum256(enc_data)
			signature, err := rsa.SignPKCS1v15(rand.Reader, &private_key, crypto.SHA256, hashed[:])
			if err != nil {
				panic(err)
			}
			signature_str := string(signature)
			log.Printf("signature created %s", signature_str)
			// Store it
			addEncData(account_id, fak, EncryptedCore{
				Id:               s.ActionID,
				EncAES:           encryptedAES,
				EncData:          enc_data,
				EncDataSignature: signature,
			})
		}
		PrintJson(fakHolder)
	}

	// Testing decrypt here - Should remove this code once integration is done
	sample_encryption := fakHolder[account_id][fak].Encryptions[0]
	sample_pb_key := shell_data.Resources[0].RSAPublicKey  // Role key
	sample_pv_key := shell_data.Resources[0].RSAPrivateKey // Role key
	// Enc AES with role public key and pass for testing
	accountPrivateKey := private_key
	decryptedAES, decryptAESErr := rsa.DecryptPKCS1v15(rand.Reader, &accountPrivateKey, sample_encryption.EncAES)
	if decryptAESErr != nil {
		panic(decryptAESErr)
	}
	encryptedAES, encrpytionAESErr := rsa.EncryptPKCS1v15(rand.Reader, &sample_pb_key, decryptedAES)
	if encrpytionAESErr != nil {
		fmt.Printf("Error encrypting AES key: %s\n", encrpytionAESErr.Error())
		os.Exit(1)
	}
	log.Printf("encrypted the AES for the role key")
	decrypted_data := decryptData(account_id, fak, "exec_hello", encryptedAES, sample_pv_key)
	log.Printf("decrypted_data =>>> ", string(decrypted_data))

	// Test FAK Signature validation
	sign_result := validateFAKSignature(account_id, []byte(fak), getFAKHash(account_id, fak), accountPublicKey)
	log.Printf("FAK sign_result =>>> %s", strconv.FormatBool(sign_result))
}

func getEncrpytedDataLength(account_id string, fak string) int {
	return len(fakHolder[account_id][fak].Encryptions)
}

func decryptData(account_id string, fak string, action_id string, enc_aes []byte, role_private rsa.PrivateKey) []byte {
	/*
		Todo: combine last 3 params and accept serialized obj
	*/
	decryptedAES, decryptAESErr := rsa.DecryptPKCS1v15(rand.Reader, &role_private, enc_aes)
	if decryptAESErr != nil {
		panic(decryptAESErr)
	}
	block, err := aes.NewCipher(decryptedAES)
	if err != nil {
		fmt.Printf("Error reading key: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Key: %s\n", hex.EncodeToString(decryptedAES))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("Error initializing AEAD: %s\n", err.Error())
		os.Exit(1)
	}
	nonceSize := gcm.NonceSize()
	encryption_list := fakHolder[account_id][fak].Encryptions
	resource_values := func(enc_obj EncryptedCore) bool { return enc_obj.Id == action_id }
	filtered_value := filterEncryptedData(encryption_list, resource_values)
	if len(filtered_value) > 0 {
		encrypted_obj := filtered_value[0]
		enc_value := encrypted_obj.EncData
		if len(enc_value) < nonceSize {
			fmt.Printf("Ciphertext too short.")
		}
		nonce_n := enc_value[0:nonceSize]
		msg := enc_value[nonceSize:]
		original_value, openErr := gcm.Open(nil, nonce_n, msg, nil)
		if openErr != nil {
			fmt.Printf("Open Err: %s\n", err.Error())
		}
		log.Printf("decrypted value: %s", string(original_value))
		return original_value
	}
	return []byte("Decrypt Error!")
}

func getFAKHash(account_id string, fak string) [32]byte {
	return fakHolder[account_id][fak].DataHash
}

func validateFAKSignature(account_id string, signature []byte, fak_hash [32]byte, public_key rsa.PublicKey) bool {
	err := rsa.VerifyPKCS1v15(&public_key, crypto.SHA256, fak_hash[:], signature)
	if err != nil {
		fmt.Printf("validateFAKSignature Err: %s\n", err.Error())
		return false
	}
	log.Printf("Verified Signature!")
	return true
}

/*
	suite := edwards25519.NewBlakeSHA256Ed25519()
	x, y, z := ElGamalEncrypt(suite, role_kyber_public_key, []byte(s.Code))
	log.Printf("encrypted the AES with role key - %s", filtered_value[0].R_type)
	log.Printf("XXX ======================================== XXX")
	PrintJson(x)
	PrintJson(y)
	PrintJson(z)
	log.Printf("XXX ======================================== XXX")
	dec_val, errDecrypting := ElGamalDecrypt(suite, role_kyber_pvt_key, x, y)
	log.Printf(">>> ======================================== <<<")
	if errDecrypting != nil {
		panic(errDecrypting)
	}
	log.Printf(string(dec_val))
	log.Printf(">>> ======================================== <<<")
*/

// Use these functions to perform asymmetric encrpytion
func ElGamalEncrypt(group kyber.Group, pubkey kyber.Point, message []byte) (
	K, C kyber.Point, remainder []byte) {

	// Embed the message (or as much of it as will fit) into a curve point.
	M := group.Point().Embed(message, random.New())
	max := group.Point().EmbedLen()
	if max > len(message) {
		max = len(message)
	}
	remainder = message[max:]
	// ElGamal-encrypt the point to produce ciphertext (K,C).
	k := group.Scalar().Pick(random.New()) // ephemeral private key
	K = group.Point().Mul(k, nil)          // ephemeral DH public key
	S := group.Point().Mul(k, pubkey)      // ephemeral DH shared secret
	C = S.Add(S, M)                        // message blinded with secret
	return
}

func ElGamalDecrypt(group kyber.Group, prikey kyber.Scalar, K, C kyber.Point) (
	message []byte, err error) {

	// ElGamal-decrypt the ciphertext (K,C) to reproduce the message.
	S := group.Point().Mul(prikey, K) // regenerate shared secret
	M := group.Point().Sub(C, S)      // use to un-blind the message
	message, err = M.Data()           // extract the embedded data
	return
}
