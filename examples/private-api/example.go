package privateapi

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/v2crypto/aevo-go-sdk/aevo"
)

func ExamplePrivateAPI() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	client, err := aevo.NewClient("testnet", os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Println(err)
	}

	orders, err := client.GetOrders()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(orders))
}
