package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func PrintJson(object any) {

	b, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

func filterShellData(ss []ResourceKey, test func(ResourceKey) bool) (ret []ResourceKey) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterEncryptedData(ss []EncryptedCore, test func(EncryptedCore) bool) (ret []EncryptedCore) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func randBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}
