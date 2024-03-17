package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type (
	AssetPrice struct {
		Timestamp int64   `json:"timestamp,string"`
		Price     float64 `json:"price,string"`
	}
	AssetPriceHistory struct {
		AssetHistory [][]string `json:"history"`
	}
	CoingeckoStatistic struct {
		TickerID                 string `json:"ticker_id"`
		BaseCurrency             string `json:"base_currency"`
		TargetCurrency           string `json:"target_currency"`
		TargetVolume             string `json:"target_volume"`
		ProductType              string `json:"product_type"`
		OpenInterest             string `json:"open_interest"`
		IndexPrice               string `json:"index_price"`
		IndexCurrency            string `json:"index_currency"`
		NextFundingRateTimestamp string `json:"next_funding_rate_timestamp"`
		FundingRate              string `json:"funding_rate"`
		ContractType             string `json:"contract_type"`
		ContractPriceCurrency    string `json:"contract_price_currency"`
	}
	Funding struct {
		NextEpoch   string `json:"next_epoch"`
		FundingRate string `json:"funding_rate"`
	}
	AevoTime struct {
		Name      string `json:"name"`
		Timestamp int    `json:"timestamp"`
		Time      string `json:"time"`
		Sequence  int    `json:"sequence"`
		Block     int    `json:"block"`
	}
	AevoSignedOrder struct {
		IsBuy      bool           `json:"is_buy"`
		Instrument int            `json:"instrument"`
		LimitPrice string         `json:"limit_price"`
		Amount     int64          `json:"amount,string"`
		Timestamp  int64          `json:"timestamp,string"`
		Salt       int64          `json:"salt,string"`
		Maker      common.Address 		  `json:"maker"`
		Signature  string         `json:"signature"`
		PostOnly 	bool	      `json:"post_only"`
		ReduceOnly 	bool          `json:"reduce_only"`
		TimeInForce string 		  `json:"time_in_force"`
	}

	OrderBook struct {
		Bids [][]string
		Asks [][]string
		Error string `json:"error"`
	}

	CreateOrderRes struct {
		OrderID          string `json:"order_id"`
		Account          string `json:"account"`
		InstrumentName   string `json:"instrument_name"`
		InstrumentID     string `json:"instrument_id"`
		InstrumentType   string `json:"instrument_type"`
		OrderType        string `json:"order_type"`
		OrderStatus      string `json:"order_status"`
		Side             string `json:"side"`
		Amount           string `json:"amount"`
		Price            string `json:"price"`
		Filled           string `json:"filled"`
		CreatedTimestamp string `json:"created_timestamp"`
		Timestamp        string `json:"timestamp"`
		SystemType       string `json:"system_type"`
		InitialMargin    string `json:"initial_margin"`
		AvgPrice         string `json:"avg_price"`
		Error 			 string `json:"error"`
	}
)
