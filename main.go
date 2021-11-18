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
	sender := GetBalance(*client, "0x02CE4C4C46fA0186A2aa98757CC20E674780c361")
	fmt.Println(sender)
	a2 := GetBalance(*client, "0x146CC8072e211C123Fd2da06148A445F0349C8ef")
	fmt.Println(a2)
	a1 := GetBalance(*client, "0x6C2D771683796F30F7c20DAFD4feb5560fddD185")
	fmt.Println(a1)
}
