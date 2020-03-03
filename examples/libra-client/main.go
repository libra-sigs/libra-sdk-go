package main

import (
	"fmt"
	"time"

	libra "github.com/libra-sigs/libra-sdk-go/libra"
)

func main() {
	client, err := libra.NewClient("ac.testnet.libra.org:8000", time.Second*3)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	account := "5f0e5ac8adf42bb02a13c53b69d6e90f0e5ec020ddbe5c2e7acb70a9d48a85ce"
	accountState, err := client.GetAccountState(account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Account resource: %v\n", accountState.AccountResource)

	translation, err := client.GetAccountTransactionBySequenceNumber(account, 0, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Translation : %v\n", translation)
}
