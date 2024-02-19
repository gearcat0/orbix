package github.com/gearcat0/orbix-sdk/order

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"pro-go-sdk/pkg/client"
	"pro-go-sdk/pkg/signature"
)

type Option struct {
	Type   string `json:"type"` // possible value ['limit','market']
	Pair   string `json:"pair"`
	Side   string `json:"side"` // possible value ['buy','sell']
	Price  string `json:"price,omitempty"`
	Amount string `json:"amount,omitempty"`
	Nonce  int64  `json:"nonce"`
}



func Create(c client.Client, option Option) (*Order, error) {
	bb, err := json.Marshal(option)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	err = json.Unmarshal(bb, &params)
	if err != nil {

	}
	sig := signature.Sign(c.APISecret(), params)
	order := &Order{}
	_, err = resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetBody(option).
		SetResult(order).
		Post(c.URL() + "/orders/")
	return order, err
}
