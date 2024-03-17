package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/v2crypto/aevo-go-sdk/aevo"
	"github.com/v2crypto/aevo-go-sdk/models"
)

func main() {
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
	getInstrument, err := client.GetInstrumentByName("ETH-PERP")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(getInstrument))
	body, err := client.CreateAndSignOrder(models.AevoSignedOrder{
		Instrument: 2054,
		IsBuy:      true,
		Amount:     1,
		LimitPrice: "2500",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
}
