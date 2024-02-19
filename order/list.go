package github.com/gearcat0/orbix-sdk/order

import (
	"gopkg.in/resty.v1"
	"pro-go-sdk/pkg/client"
	"pro-go-sdk/pkg/signature"
)


type ListResp []Order

func List(c client.Client, params ...string) (ListResp, error) {
	sig := signature.Sign(c.APISecret(), nil)
	pair := ""
	status := "open" // default status for search
	side := "buy"
	offset := "0"
	if len(params) == 1 {
		pair = params[0]
	}

	if len(params) == 2 {
		pair = params[0]
		status = params[1] // param 2 is status for search
	}
	if len(params) == 3 {
		pair = params[0]
		status = params[1]
		side = params[2]
	}
	if len(params) == 4 {
		pair = params[0]
		status = params[1]
		side = params[2]
		offset = params[3]
	}

	orders := ListResp{}
	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetQueryParams(map[string]string{
			"pair":   pair,
			"limit":  "50",
			"offset": offset,
			"status": status,
			"side":   side,
		})

	req.SetResult(&orders)
	_, err := req.Get(c.URL() + "/orders/user")

	if err != nil {
		return nil, err
	}

	return orders, nil
}
