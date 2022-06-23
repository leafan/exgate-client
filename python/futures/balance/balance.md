
## FutureTrade_FutureBalance

<a id="opIdFutureTrade_FutureBalance"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/balance', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/balance", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/balance`

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


### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int32)|false|none|none|
|message|string|false|none|none|
|details|[[protobufAny](#schemaprotobufany)]|false|none|none|

<h2 id="tocS_tradeBalanceResponse">tradeBalanceResponse</h2>
<!-- backwards compatibility -->
<a id="schematradebalanceresponse"></a>
<a id="schema_tradeBalanceResponse"></a>
<a id="tocStradebalanceresponse"></a>
<a id="tocstradebalanceresponse"></a>

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