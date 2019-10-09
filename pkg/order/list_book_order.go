package order

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"pro-go-sdk/pkg/client"
)

type listOrdersResponse struct {
	Bids []Order `json:"bid"`
	Asks []Order `json:"ask"`
}

func ListOrdersBook(c client.Client, pairing string) ([]Order, []Order, error) {
	host := c.URL()
	result := &listOrdersResponse{}
	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"Symbol": pairing,
		}).
		SetHeader("Accept", "application/json").
		SetResult(&listOrdersResponse{}).
		Get(fmt.Sprintf("%s/orders/?pair=%s", host, pairing))
	_ = json.Unmarshal(resp.Body(), result)
	return result.Bids, result.Asks, err
}
