//test rsa encrypt
package main

import(
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
)


func main2(){

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	publicKey := &privateKey.PublicKey


	privateKeyDos, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	publicKeyDos := &privateKeyDos.PublicKey

	fmt.Println("Private key: ", privateKey)
	fmt.Println("Public key: ", publicKey)
	fmt.Println("Private key: ", privateKeyDos)
	fmt.Println("Private key: ", publicKeyDos)


}