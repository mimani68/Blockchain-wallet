package main

import (
	"fmt"

	bip39 "github.com/tyler-smith/go-bip39"
)

func mnemonicGenerator() (*string, error) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	return &mnemonic, nil
}

func main() {
	a, _ := mnemonicGenerator()
	fmt.Println(*a)
}
