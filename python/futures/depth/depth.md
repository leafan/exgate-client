
## Universal_Depth

<a id="opIdUniversal_Depth"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/get_depth/{symbol}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/get_depth/{symbol}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/get_depth/{symbol}`

<h3 id="universal_depth-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所|
|symbol|path|string|true|币种|
|limit|query|string(int64)|false|默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000].|
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

<h3 id="universal_depth-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1universalDepthResponse](#schemav1universaldepthresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>


## Universal_PullDepthStreams

<a id="opIdUniversal_PullDepthStreams"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/pull_depth/{symbol}/{level}/{push_rate}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/pull_depth/{symbol}/{level}/{push_rate}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/pull_depth/{symbol}/{level}/{push_rate}`

*拉取有限档深度信息,走websocket*

<h3 id="universal_pulldepthstreams-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|path|string|true|交易对名称|
|level|path|string|true|level表示几档买卖单信息, 可选 5/10/20档|
|push_rate|path|string|true|push Speed: 250ms 或 500ms 或 100ms|

> Example responses

> 200 Response

```json
{
  "desc": "string",
  "status": "string"
}
```

<h3 id="universal_pulldepthstreams-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[universalPullDepthResponse](#schemauniversalpulldepthresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>