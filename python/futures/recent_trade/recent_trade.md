
## FutureTrade_FutureRecentTradesList

<a id="opIdFutureTrade_FutureRecentTradesList"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/recent_trades', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/recent_trades", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/recent_trades`

<h3 id="futuretrade_futurerecenttradeslist-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|query|string|false|required 币种.|
|limit|query|string(int64)|false|required 分页页码.|
|isLocal|query|boolean|false|本地.|

> Example responses

> 200 Response

```json
{
  "trades": [
    {
      "id": 0,
      "price": "string",
      "qty": "string",
      "quoteQty": "string",
      "time": 0,
      "isBuyerMaker": true
    }
  ]
}
```

<h3 id="futuretrade_futurerecenttradeslist-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[tradeRecentTradesListResponse](#schematraderecenttradeslistresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>
