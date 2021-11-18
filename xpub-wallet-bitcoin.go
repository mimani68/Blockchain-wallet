package main

import (
	"fmt"

	hdwallet "github.com/wemeetagain/go-hdwallet"
)

func bitCoinDepisteAddress() {
	// Generate a random 256 bit seed
	seed, err := hdwallet.GenSeed(256)
	if err != nil {
		fmt.Println(err)
	}

	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Generate new child key based on private or public key
	childprv, err := masterprv.Child(0)
	fmt.Printf("private of child [0]: %s \n", childprv)

	childpub, err := masterpub.Child(0)

	// Create bitcoin address from public key
	address := childpub.Address()
	fmt.Printf("Address: %s \n", address)

	var i uint32
	for i = 0; i < 10; i++ {
		// Convenience string -> string Child and ToAddress functions
		walletstring := childpub.String()
		childstring, _ := hdwallet.StringChild(walletstring, i)
		childaddress, _ := hdwallet.StringAddress(childstring)
		fmt.Printf("child %d Address: %s \n", i, childaddress)
	}
}
