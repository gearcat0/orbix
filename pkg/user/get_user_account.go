package user


import (
	"gopkg.in/resty.v1"
	"fmt"
	"pro-go-sdk/pkg/client"
	"pro-go-sdk/pkg/signature"
)

func Get(c client.Client) (*UserAccount, error) {
	sig := signature.Sign(c.APISecret(), nil)
	result := &UserAccount{}

	_, err := resty.R().
		SetHeader("Authorization", c.APIKey()).
		SetHeader("Signature", sig).
		SetResult(result).
		Get(c.URL() + "/users/" + c.UserID())

	if err != nil {
		fmt.Printf("Get balance error = %+v", err)
	}

	return result, nil
}
