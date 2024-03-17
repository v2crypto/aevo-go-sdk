package aevo

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/v2crypto/aevo-go-sdk/models"
)

func TestAevoPrivate(t *testing.T) {
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

	res, err := client.GetAccount()
	assert.Nil(t, err)
	fmt.Println(res)
}