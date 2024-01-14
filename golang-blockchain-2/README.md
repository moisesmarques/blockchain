# blockchain-go-18
The blockchain repo for Go Group 28


# wallet
124vDn5HpDaXz87vKaapfv4KFYTzjsRHfZ
18dZ1XaH27CkjnZXXH95kdfmS664KJv1Rg

# Create FAK
`go run . createfak -address 1P7ntZWZYiNNsFHTUUcJ28hCY8ZpqxDu6V`

# Encryption and decryption

By default, the data that must be shared across shards are encrypted as soon as the shell creation is performed. The action data or code is encrypted by an AES key and stored against FAK with their action ids so that they can be retrieved easily whenever required.

Since the AES key is used as an intermediary key to encrypt and decrypt the data for more than two parties to access, it's known as the symmetric category.

The encryption process requires palette data and shell creation data, to begin with. The action-related code (data) will be taken from the palette and shard keys needed for encryption will be picked from the shell creation. 

The data would be encrypted by the AES key and shared across the network. The AES key will also be shared with the shards after being encrypted with the shard’s public key. 

# Decryption
The owner entities who have access to the code of related action_id will decrypt their AES key and try to decrypt the data. 

FAK validators will have functions exposed to facilitate these operations. The functions that are available today are 

`func decryptData(account_id string, fak string, action_id string, enc_aes []byte, role_private rsa.PrivateKey)` returns `[]byte` encrpyted data

`func getFAKHash(account_id string, fak string)` returns `[32]byte` FAK Hash

`func validateFAKSignature(account_id string, signature []byte, fak_hash [32]byte, public_key rsa.PublicKey)` returns `bool` true/false based on the signature’s validation results.

# References
https://lucid.app/lucidchart/17f15de8-c02c-4d77-b945-7adf0ef75ff5/edit?invitationId=inv_02152011-9a47-4d10-9c15-8bb124d766dc&page=gEUp1W2Bcs.~#