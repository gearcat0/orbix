package order

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"github.com/gearcat0/orbix/client"
	"github.com/gearcat0/orbix/signature"
)

type CancelRequest struct {
	Pair string `json:"pair"`
}


func Cancel(c client.Client, orderID int64, pair string) error {
	cancelRequest := &CancelRequest{
		Pair: pair,
	}
	bb, err := json.Marshal(cancelRequest)
	if err != nil {
		return err
	}

	params := map[string]interface{}{}
	err = json.Unmarshal(bb, &params)
	if err != nil {
		return err
	}

	sig := signature.Sign(c.APISecret(), params)
	_, err = resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetHeader("Content-Type", "application/json").
		SetBody(cancelRequest).
		Delete(fmt.Sprintf("%s/orders/%d", c.URL(), orderID))
	return err
}
