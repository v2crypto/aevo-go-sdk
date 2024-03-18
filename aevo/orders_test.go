package aevo

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/v2crypto/aevo-go-sdk/models"
)

func TestAevoOrder(t *testing.T) {
	err := godotenv.Load()  
	if err != nil {  
		panic(err)  
	}

	client, err := NewClient(models.ClientOption{
		ChainType: "mainnet",
		Address: os.Getenv("ADDRESS"),
		ApiKey: os.Getenv("API_KEY"),
		ApiSecret: os.Getenv("API_SECRET"),
		SigningKey: os.Getenv("SIGNING_KEY"),
	})
	if err != nil {
		fmt.Println(err)
	}

	t.Run("GetInstrumentByName", func(t *testing.T) {
		res, err := client.GetInstrumentByName("ETH-PERP")
		assert.Nil(t, err)
		fmt.Println(string(res))
	})


	t.Run("GetOrderBook", func(t *testing.T) {
		orderBook, err := client.GetOrderBook("ETH-PERP")
		assert.Nil(t, err)
		fmt.Println(orderBook)
	})

	t.Run("MarketOrder", func(t *testing.T) {
		res, err := client.MarketOrder(false, 1, 0.01)
		assert.Nil(t, err)
		fmt.Println(res)
	})

	t.Run("GetTradeHistory", func(t *testing.T) {
		res, err := client.GetTradeHistory(models.TradeHistoryReq{
			StartTime: 0,
			InstrumentType: "PERPETUAL",
			Limit: 5,
			TradeTypes: "trade_types=trade&trade_types=liquidation&trade_types=settlement",
		})
		assert.Nil(t, err)
		fmt.Println(res)
	})




}