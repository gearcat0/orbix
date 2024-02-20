package user

type WalletCoin struct {
	Addresses          []string `json:"addresses"`
	AvailableBalance string   `json:"available_balance"`
}

type Wallets struct {
	Btc  WalletCoin `json:"btc"`
	Btg  WalletCoin `json:"btg"`
	Cpy  WalletCoin `json:"cpy"`
	Elec WalletCoin `json:"elec"`
	Eth  WalletCoin `json:"eth"`
	Etp  WalletCoin `json:"etp"`
	Jfin WalletCoin `json:"jfin"`
	Knc  WalletCoin `json:"knc"`
	Ltc  WalletCoin `json:"ltc"`
	Neo  WalletCoin `json:"neo"`
	Omg  WalletCoin `json:"omg"`
	Powr WalletCoin `json:"powr"`
	Rpx  WalletCoin `json:"rpx"`
	Thb  WalletCoin `json:"thb"`
	Usdt WalletCoin `json:"usdt"`
	Wax  WalletCoin `json:"wax"`
	Xlm  WalletCoin `json:"xlm"`
	Xrp  WalletCoin `json:"xrp"`
	Xzc  WalletCoin `json:"xzc"`
}

type Limit struct {
	Available string `json:"available"`
	Capacity  string `json:"capacity"`
}

type Limits struct {
	Btc  Limit `json:"btc"`
	Btg  Limit `json:"btg"`
	Cpy  Limit `json:"cpy"`
	Elec Limit `json:"elec"`
	Eth  Limit `json:"eth"`
	Etp  Limit `json:"etp"`
	Jfin Limit `json:"jfin"`
	Knc  Limit `json:"knc"`
	Ltc  Limit `json:"ltc"`
	Neo  Limit `json:"neo"`
	Omg  Limit `json:"omg"`
	Powr Limit `json:"powr"`
	Rpx  Limit `json:"rpx"`
	Thb  Limit `json:"thb"`
	Usdt Limit `json:"usdt"`
	Wax  Limit `json:"wax"`
	Xlm  Limit `json:"xlm"`
	Xrp  Limit `json:"xrp"`
	Xzc  Limit `json:"xzc"`
}

type UserAccount struct {
	ID                        int64     `json:"id"`
	IdentityVerificationLevel string    `json:"identity_verification_level"`
	UserWallets               Wallets `json:"wallets"`
	UserLimits                Limits  `json:"limits"`
}
