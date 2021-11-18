package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func createConnection() {
	var url = os.Getenv("PROVIDER_URL")
	var err error
	client, err = ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	} else {
		fmt.Println("Sucess! you are connected to the Ethereum Network")
	}
}
