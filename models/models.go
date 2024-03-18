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

	ClientOption struct {
		ChainType  string
		Address string
		ApiKey string
		ApiSecret string
		SigningKey string
	}

	Account struct {
		Error 		string `json:"error"`
		Account     string `json:"account"`
		Username    string `json:"username"`
		AccountType string `json:"account_type"`
		Portfolio   bool   `json:"portfolio"`
		Equity      string `json:"equity"`
		Balance     string `json:"balance"`
		Credit      string `json:"credit"`
		Credited    bool   `json:"credited"`
		Collaterals []struct {
			CollateralAsset        string `json:"collateral_asset"`
			Balance                string `json:"balance"`
			AvailableBalance       string `json:"available_balance"`
			WithdrawableBalance    string `json:"withdrawable_balance"`
			MarginValue            string `json:"margin_value"`
			CollateralValue        string `json:"collateral_value"`
			CollateralYieldBearing bool   `json:"collateral_yield_bearing"`
		} `json:"collaterals"`
		AvailableBalance  string  `json:"available_balance"`
		InitialMargin     string  `json:"initial_margin"`
		MaintenanceMargin string  `json:"maintenance_margin"`
		EmailAddress      string  `json:"email_address"`
		InLiquidation     bool    `json:"in_liquidation"`
		ReferralBonus     float64 `json:"referral_bonus"`
		RefereeDiscount   float64 `json:"referee_discount"`
		HasBeenReferred   bool    `json:"has_been_referred"`
		IntercomHash      string  `json:"intercom_hash"`
		Positions         []struct {
			Asset             string `json:"asset"`
			InstrumentID      string `json:"instrument_id"`
			InstrumentName    string `json:"instrument_name"`
			InstrumentType    string `json:"instrument_type"`
			Amount            string `json:"amount"`
			Side              string `json:"side"`
			MarkPrice         string `json:"mark_price"`
			AvgEntryPrice     string `json:"avg_entry_price"`
			UnrealizedPnl     string `json:"unrealized_pnl"`
			MaintenanceMargin string `json:"maintenance_margin"`
			LiquidationPrice  string `json:"liquidation_price"`
			MarginType        string `json:"margin_type"`
			Triggers          struct {
			} `json:"triggers"`
		} `json:"positions"`
		SigningKeys []struct {
			SigningKey       string `json:"signing_key"`
			Expiry           string `json:"expiry"`
			CreatedTimestamp string `json:"created_timestamp"`
		} `json:"signing_keys"`
		APIKeys []struct {
			Name             string `json:"name,omitempty"`
			APIKey           string `json:"api_key"`
			ReadOnly         bool   `json:"read_only"`
			CreatedTimestamp string `json:"created_timestamp"`
		} `json:"api_keys"`
		FeeStructures []struct {
			Asset          string `json:"asset"`
			InstrumentType string `json:"instrument_type"`
			TakerFee       string `json:"taker_fee"`
			MakerFee       string `json:"maker_fee"`
		} `json:"fee_structures"`
		Leverages []struct {
			InstrumentID string `json:"instrument_id"`
			Leverage     string `json:"leverage"`
			MarginType   string `json:"margin_type"`
		} `json:"leverages"`
		Permissions []any `json:"permissions"`
	}

	TradeHistoryReq struct {
		StartTime int64
		Asset string
		EndTime int64
		TradeTypes string
		InstrumentType string
		OptionType string
		Limit int64
		Offset int64
	}

	TradeHistoryRes struct {
	Count        string `json:"count"`
	TradeHistory []struct {
		TradeID          string `json:"trade_id"`
		TradeType        string `json:"trade_type"`
		InstrumentID     string `json:"instrument_id"`
		InstrumentName   string `json:"instrument_name"`
		InstrumentType   string `json:"instrument_type"`
		Asset            string `json:"asset"`
		SpotPrice        string `json:"spot_price,omitempty"`
		Account          string `json:"account"`
		Amount           string `json:"amount"`
		Side             string `json:"side"`
		Fees             string `json:"fees"`
		CreatedTimestamp string `json:"created_timestamp"`
		Pnl              string `json:"pnl"`
		Price            string `json:"price,omitempty"`
		OrderID          string `json:"order_id,omitempty"`
		OrderType        string `json:"order_type,omitempty"`
		AggOrderID       string `json:"agg_order_id,omitempty"`
		TradeStatus      string `json:"trade_status,omitempty"`
		Liquidity        string `json:"liquidity,omitempty"`
		TakerTradeID     string `json:"taker_trade_id,omitempty"`
		FarmBoost        string `json:"farm_boost,omitempty"`
		MarkPrice        string `json:"mark_price,omitempty"`
		FeeRate          string `json:"fee_rate,omitempty"`
	} `json:"trade_history"`
}
)
