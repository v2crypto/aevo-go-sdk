package privateapi

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/v2crypto/aevo-go-sdk/aevo"
	"github.com/v2crypto/aevo-go-sdk/models"
)

func ExamplePrivateAPI() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	client, err := aevo.NewClient(models.ClientOption{
		ChainType: "mainnet",
		Address: os.Getenv("ADDRESS"),
		ApiKey: os.Getenv("API_KEY"),
		ApiSecret: os.Getenv("API_SECRET"),
		SigningKey: os.Getenv("SIGNING_KEY"),
	})
	if err != nil {
		fmt.Println(err)
	}

	orders, err := client.GetOrders()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(orders))
}
