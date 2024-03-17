package aevo

import "github.com/v2crypto/aevo-go-sdk/models"

type Client struct {
	baseUrl    string
	chainType  string
	address string
	apiKey string
	apiSecret string
	signingKey string
}

const (
	testnetUrl = "https://api-testnet.aevo.xyz/"
	mainnetUrl = "https://api.aevo.xyz/"
)

func NewClient(option models.ClientOption) (*Client, error) {
	baseUrl := testnetUrl
	if option.ChainType == "mainnet" {
		baseUrl = mainnetUrl
	}
	return &Client{
		baseUrl:    baseUrl,
		chainType:  option.ChainType,
		address: option.Address,
		apiKey: option.ApiKey,
		apiSecret: option.ApiSecret,
		signingKey: option.SigningKey,
	}, nil
}
