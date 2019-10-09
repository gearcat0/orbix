# pro-go-sdk
##SDK Features
1. [ Client ](#client)
2. [ Create. ](#create)
3. [ Cancel. ](#cencel)
4. [ List. ](#list)
5. [ List book order. ](#listborder)
6. [ User. ](#user)


<a name="client"></a>
## 1. Client
```dtd
client := client.NewClient(
    url: "https://api.tdax.com/api/",
    userID : "1"// let's say your user id is 1
    apiKey : "live-xxxx",
    apiSecret : "xxxxx"
)
```
<a name="create"></a>
## 2. Create
Call example.
```dtd
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

<a name="cencel"></a>
## 3. Cancel
Call example
```dtd
err:= Cancel(client,orderId,pair)
```


<a name="cencel"></a>
## 4. Cancel
Call example
```dtd
err:= Cancel(client,orderId,pair)
```
