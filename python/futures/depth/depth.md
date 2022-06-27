
## FutureTrade_FutureDepth

<a id="opIdFutureTrade_FutureDepth"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/depth/{symbol}/{limit}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/depth/{symbol}/{limit}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/depth/{symbol}/{limit}`

<h3 id="futuretrade_futuredepth-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所|
|symbol|path|string|true|币种|
|limit|path|string(int64)|true|数量|
|isLocal|query|boolean|false|本地.|
|tolerate_interval|query|string(int64)|false|允许延迟间隔单位:秒.|

> Example responses

> 200 Response

```json
{
  "lastUpdateId": 0,
  "timeStamp": 0,
  "bid_list": [
    {
      "bid": [
        {
          "price": "string",
          "amount": "string"
        }
      ]
    }
  ],
  "ask_list": [
    {
      "bid": [
        {
          "price": "string",
          "amount": "string"
        }
      ]
    }
  ]
}
```

<h3 id="futuretrade_futuredepth-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[tradeDepthResponse](#schematradedepthresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>