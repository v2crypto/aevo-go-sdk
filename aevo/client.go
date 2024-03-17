package aevo

type Client struct {
	baseUrl    string
	chainType  string
	address string
}

const (
	testnetUrl = "https://api-testnet.aevo.xyz/"
	mainnetUrl = "https://api.aevo.xyz/"
)

func NewClient(chainType, address string) (*Client, error) {
	baseUrl := testnetUrl
	if chainType == "mainnet" {
		baseUrl = mainnetUrl
	}
	return &Client{
		baseUrl:    baseUrl,
		chainType:  chainType,
		address: address,
	}, nil
}
