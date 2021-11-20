package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func recoverAccount(mnemonic string, accountNumber int) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	pathHD := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", accountNumber))
	account, err := wallet.Derive(pathHD, false)
	if err != nil {
		log.Fatal(err)
	}
	priv, _ := wallet.PrivateKeyHex(account)
	fmt.Printf("%s\n", priv)
}

func createDepositeAddress(mnemonic string) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		path := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		fmt.Printf("=========| %d |=======\n", i)
		pathHD := hdwallet.MustParseDerivationPath(path)
		account, err := wallet.Derive(pathHD, false)
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

func accumolativeBalanceOfWallet(mnemonic string) {
	fmt.Println("Comming soon")
}

func dev(mnemonic string, path string) {
	fmt.Printf("Path: %s \n", path)
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	pathHD := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(pathHD, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Address: %s\n", account.Address.Hex())
}
