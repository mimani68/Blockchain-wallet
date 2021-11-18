package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
)

var client *ethclient.Client

func createConnection() {
	// var url = os.Getenv("PROVIDER_URL")
	var url = "https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	var err error
	client, err = ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
	} else {
		fmt.Println("Sucess! you are connected to the Ethereum Network")
	}
}

func GetBalance(client ethclient.Client, address string) string {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance.String()
}

func main() {
	createConnection()
	balance := GetBalance(*client, "0x02CE4C4C46fA0186A2aa98757CC20E674780c361")
	fmt.Println(balance)
}
