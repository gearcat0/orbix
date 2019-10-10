# pro-go-sdk
##SDK Features

## 1. Client
```
client := client.NewClient(
    url: "https://api.tdax.com/api/",
    userID : "1"// let's say your user id is 1
    apiKey : "live-xxxx",
    apiSecret : "xxxxx"
)

```

## 2. Create
Call example.
```
options := Option{
    Type:   "limit",
    Pair:   "btc_thb",
    Side:   "sell",
    Nonce:  2731832,
    Price:  "247315.50",
    Amount: "0.649",
}
order, err := Create(client, options)

```


## 3. Cancel
Call example

```
err:= Cancel(client,orderId,pair)
```


## 4. Cancel

Call example

```
err:= Cancel(client,orderId,pair)
```

