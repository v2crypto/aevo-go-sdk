package aevo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/v2crypto/aevo-go-sdk/models"
)

// GetAccount returns the account's information including API keys, signing keys and positions.
func (c *Client) GetAccount() (*models.Account, error) {
	url := fmt.Sprintf("%s/account", c.baseUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	headers := c.getHeaders("/account", "GET", "")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var account models.Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		return &models.Account{}, err
	}
	return &account, nil
}

// GetTradeHistory returns the account's trade history.
func (c *Client) GetTradeHistory(param models.TradeHistoryReq) (*models.TradeHistoryRes, error) {
	url := fmt.Sprintf("%s/trade-history?start_time=%d&end_time=%d&asset=%s&limit=%d&offset=%d&instrument_type=%s&%s&sort=created_timestamp&sort_order=DESC",
	c.baseUrl, param.StartTime, param.EndTime, param.Asset, param.Limit, param.Offset, param.InstrumentType, param.TradeTypes)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	headers := c.getHeaders("/trade-history", "GET", "")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var tradeHistory models.TradeHistoryRes
	err = json.Unmarshal(body, &tradeHistory)
	if err != nil {
		return &models.TradeHistoryRes{}, err
	}
	return &tradeHistory, nil
}