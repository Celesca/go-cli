package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func encrypt(filePath string, key string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	for i := range data {
		data[i] ^= key[i%len(key)] // Or operator
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

func decrypt(filePath string, key string) error {
	return encrypt(filePath, key)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <encrypt/decrypt> <filepath> <key>")
		os.Exit(1)
	}

	command := os.Args[1]
	filePath := os.Args[2]
	key := os.Args[3]

	var err error
	if command == "encrypt" {
		err = encrypt(filePath, key)
	} else if command == "decrypt" {
		err = decrypt(filePath, key)
	} else {
		fmt.Println("Invalid command. Use 'encrypt' or 'decrypt'")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("File encrypted/decrpyted successfully!")

}
