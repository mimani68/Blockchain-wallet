package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// //
	// // connect to network
	// //
	// createConnection()

	//
	// Get balance of address in network
	//
	// sender := GetBalance(*client, "0x02CE4C4C46fA0186A2aa98757CC20E674780c361")
	// fmt.Println(sender)
	// a2 := GetBalance(*client, "0x146CC8072e211C123Fd2da06148A445F0349C8ef")
	// fmt.Println(a2)
	// a1 := GetBalance(*client, "0x6C2D771683796F30F7c20DAFD4feb5560fddD185")
	// fmt.Println(a1)

	// //
	// // Create old wallets
	// //
	// createOldWallet()

	// //
	// // Print random words
	// //
	// a, _ := mnemonicGenerator()
	// fmt.Println(*a)

	//
	// This password most store permamanetly, do not share this one.
	// $ go run mnemonic.go
	//
	mnemonic := "tonight capital wing name satisfy mule grace next mansion very retreat emotion"

	//
	// You can create unlimit account for each wallet
	//
	// createDepositeAddress(mnemonic)

	accMaster := "m/44'/60'/0'/1"
	dev(mnemonic, accMaster)
	accIndex0 := "m/44'/60'/0'/1/0"
	dev(mnemonic, accIndex0)
	accIndex1 := "m/44'/60'/0'/1/1"
	dev(mnemonic, accIndex1)

	//
	// With above paswrod we can restore specidif account
	//
	// recoverAccount(mnemonic, 52)

	// //
	// // Aggregate remain value in each address wallet
	// //
	// accumolativeBalanceOfWallet(mnemonic)

	// //
	// // BitCoin deposite address generator
	// //
	// bitCoinDepisteAddress()
}
