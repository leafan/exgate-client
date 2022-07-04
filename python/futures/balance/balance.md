
<h1 id="api-exgate-v1-futures-trade-trade-proto">api/exgate/v1/futures/trade/trade.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="api-exgate-v1-futures-trade-trade-proto-futuretrade">FutureTrade</h1>

## FutureTrade_FutureBalance

<a id="opIdFutureTrade_FutureBalance"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/balance', headers = headers)

print(r.json())

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/v1/{exchange}/balance", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/balance`

<h3 id="futuretrade_futurebalance-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|

> Example responses

> 200 Response

```json
{
  "balances": [
    {
      "accountAlias": "string",
      "asset": "string",
      "balance": "string",
      "crossWalletBalance": "string",
      "crossUnPnl": "string",
      "availableBalance": "string",
      "maxWithdrawAmount": "string",
      "marginAvailable": true,
      "updateTime": 0
    }
  ]
}
```

<h3 id="futuretrade_futurebalance-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[tradeBalanceResponse](#schematradebalanceresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>