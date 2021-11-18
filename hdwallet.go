package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	mnemonic := "tonight capital wing name satisfy mule grace next mansion very retreat emotion"

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 20; i++ {
		a := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		fmt.Printf("=========| %d |=======\n", i)
		path := hdwallet.MustParseDerivationPath(a)
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Account address: %s\n", account.Address.Hex())

		privateKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Private key in hex: %s\n", privateKey)

		publicKey, _ := wallet.PublicKeyHex(account)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Public key in hex: %s\n", publicKey)
	}

}
