
## Universal_RecentTradesList

<a id="opIdUniversal_RecentTradesList"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/exgate/v1/{exchange}/get_recent_trades/{symbol}', headers = headers)

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
    req, err := http.NewRequest("GET", "/exgate/v1/{exchange}/get_recent_trades/{symbol}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /exgate/v1/{exchange}/get_recent_trades/{symbol}`

<h3 id="universal_recenttradeslist-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种|
|limit|query|string(int64)|false|默认:500，最大1000.|
|isLocal|query|boolean|false|本地.|
|tolerate_interval|query|string(int64)|false|允许延迟间隔单位:秒.|

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

<h3 id="universal_recenttradeslist-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[universalRecentTradesListResponse](#schemauniversalrecenttradeslistresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>