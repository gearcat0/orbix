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
	req := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetHeader("Content-Type", "application/json").
		SetBody(cancelRequest)

	resp, err := req.Delete(fmt.Sprintf("%s/orders/%d", c.URL(), orderID))
	if err != nil {
		fmt.Println("!!! FAILED TO EXECUTE REQUEST")
		return err
	}

	if resp.IsSuccess() != true {
		fmt.Println("!!! IS SUCCESS IS FALSE")
	}

	// Does this ever return a body?
	fmt.Printf("XXX %d: %s", resp.StatusCode(), resp.Body())

	return err
}
