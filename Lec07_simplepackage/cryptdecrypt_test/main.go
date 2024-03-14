package main

import (
	"fmt"

	cryptdecrpyt "github.com/Celesca/cryptdecrypt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(cryptdecrpyt.Encrypt("Hello, playground", "borntodev"))
}
