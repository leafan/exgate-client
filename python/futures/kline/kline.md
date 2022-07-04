---
title: api/exgate/v1/futures/kline/kline.proto version not set
language_tabs:
  - python: Python
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="api-exgate-v1-futures-kline-kline-proto">api/exgate/v1/futures/kline/kline.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="api-exgate-v1-futures-kline-kline-proto-kline">Kline</h1>

## Kline_PullAggtrade

<a id="opIdKline_PullAggtrade"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/aggtrade/pull/{symbol}/{interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/aggtrade/pull/{symbol}/{interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/aggtrade/pull/{symbol}/{interval}`

*拉取聚合交易数据，走websocket*

<h3 id="kline_pullaggtrade-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|path|string|true|交易对名称|
|interval|path|string|true|按照时间段拉取数据|

> Example responses

> 200 Response

```json
{
  "desc": "string",
  "status": "string"
}
```

<h3 id="kline_pullaggtrade-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineGetAggtradeResponse](#schemaklinegetaggtraderesponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_AggTradeScannerAllKLine

<a id="opIdKline_AggTradeScannerAllKLine"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/aggtrade/scanner/allkline/{cron_interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/aggtrade/scanner/allkline/{cron_interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/aggtrade/scanner/allkline/{cron_interval}`

*gen 所有交易所建立ws连接的k线数据*

<h3 id="kline_aggtradescannerallkline-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|cron_interval|path|string|true|调度时间间隔|

> Example responses

> 200 Response

```json
{
  "status": "string",
  "desc": "string"
}
```

<h3 id="kline_aggtradescannerallkline-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineAggtradeSannerKLineResp](#schemaklineaggtradesannerklineresp)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_AggTradeScannerKLine

<a id="opIdKline_AggTradeScannerKLine"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/aggtrade/scanner/kline/{symbol}/{interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/aggtrade/scanner/kline/{symbol}/{interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/aggtrade/scanner/kline/{symbol}/{interval}`

*从aggtrade中生成kline*

<h3 id="kline_aggtradescannerkline-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|path|string|true|交易对名称|
|interval|path|string|true|k线类型|

> Example responses

> 200 Response

```json
{
  "status": "string",
  "desc": "string"
}
```

<h3 id="kline_aggtradescannerkline-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineAggtradeSannerKLineResp](#schemaklineaggtradesannerklineresp)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_StopPullAggtrade

<a id="opIdKline_StopPullAggtrade"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/aggtrade/stop/{symbol}/{interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/aggtrade/stop/{symbol}/{interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/aggtrade/stop/{symbol}/{interval}`

*停止拉取聚合交易数据*

<h3 id="kline_stoppullaggtrade-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|path|string|true|交易对名称|
|interval|path|string|true|按照时间段拉取数据|

> Example responses

> 200 Response

```json
{
  "status": "string",
  "desc": "string"
}
```

<h3 id="kline_stoppullaggtrade-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineStopPullAggtradeResponse](#schemaklinestoppullaggtraderesponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_GetKlineByExgate

<a id="opIdKline_GetKlineByExgate"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/exgate/kline/{symbol}/{interval}/{startTime}/{endTime}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/exgate/kline/{symbol}/{interval}/{startTime}/{endTime}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/exgate/kline/{symbol}/{interval}/{startTime}/{endTime}`

*by local*

<h3 id="kline_getklinebyexgate-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种名称|
|interval|path|string|true|required 间隔|
|startTime|path|string(int64)|true|required 开始时间|
|endTime|path|string(int64)|true|required 结束时间|
|limit|query|string(int64)|false|required 数量.|

> Example responses

> 200 Response

```json
{
  "kline": [
    {
      "open_time": 0,
      "open": "string",
      "high": "string",
      "low": "string",
      "close": "string",
      "volume": "string",
      "close_time": 0
    }
  ]
}
```

<h3 id="kline_getklinebyexgate-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineGetKlineResponse](#schemaklinegetklineresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_GenKline

<a id="opIdKline_GenKline"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/genkline/{symbol}/{interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/genkline/{symbol}/{interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/genkline/{symbol}/{interval}`

*genkline*

<h3 id="kline_genkline-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种名称|
|interval|path|string|true|required 间隔|
|startTime|query|string(int64)|false|required 开始时间.|
|endTime|query|string(int64)|false|required 结束时间.|
|limit|query|string(int64)|false|required 数量.|

> Example responses

> 200 Response

```json
{
  "kline": [
    {
      "open_time": 0,
      "open": "string",
      "high": "string",
      "low": "string",
      "close": "string",
      "volume": "string",
      "close_time": 0
    }
  ]
}
```

<h3 id="kline_genkline-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineGetKlineResponse](#schemaklinegetklineresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_GetKline

<a id="opIdKline_GetKline"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/kline/{symbol}/{interval}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/kline/{symbol}/{interval}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/kline/{symbol}/{interval}`

*by proxy*

<h3 id="kline_getkline-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种名称|
|interval|path|string|true|required 间隔|
|startTime|query|string(int64)|false|required 开始时间.|
|endTime|query|string(int64)|false|required 结束时间.|
|limit|query|string(int64)|false|required 数量.|

> Example responses

> 200 Response

```json
{
  "kline": [
    {
      "open_time": 0,
      "open": "string",
      "high": "string",
      "low": "string",
      "close": "string",
      "volume": "string",
      "close_time": 0
    }
  ]
}
```

<h3 id="kline_getkline-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineGetKlineResponse](#schemaklinegetklineresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Kline_GetTradeRecent

<a id="opIdKline_GetTradeRecent"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/recent/{symbol}', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/recent/{symbol}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/recent/{symbol}`

<h3 id="kline_gettraderecent-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种名称|

> Example responses

> 200 Response

```json
{
  "trade": [
    {
      "id": "string",
      "time": "string",
      "price": "string",
      "amount": "string",
      "isBuyerMaker": true
    }
  ]
}
```

<h3 id="kline_gettraderecent-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[klineGetTradeRecentResponse](#schemaklinegettraderecentresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_GetKlineResponseKline">GetKlineResponseKline</h2>
<!-- backwards compatibility -->
<a id="schemagetklineresponsekline"></a>
<a id="schema_GetKlineResponseKline"></a>
<a id="tocSgetklineresponsekline"></a>
<a id="tocsgetklineresponsekline"></a>

```json
{
  "open_time": 0,
  "open": "string",
  "high": "string",
  "low": "string",
  "close": "string",
  "volume": "string",
  "close_time": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|open_time|number(double)|false|none|none|
|open|string|false|none|none|
|high|string|false|none|none|
|low|string|false|none|none|
|close|string|false|none|none|
|volume|string|false|none|none|
|close_time|number(double)|false|none|none|

<h2 id="tocS_GetTradeRecentResponseTrade">GetTradeRecentResponseTrade</h2>
<!-- backwards compatibility -->
<a id="schemagettraderecentresponsetrade"></a>
<a id="schema_GetTradeRecentResponseTrade"></a>
<a id="tocSgettraderecentresponsetrade"></a>
<a id="tocsgettraderecentresponsetrade"></a>

```json
{
  "id": "string",
  "time": "string",
  "price": "string",
  "amount": "string",
  "isBuyerMaker": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string(int64)|false|none|none|
|time|string(int64)|false|none|none|
|price|string|false|none|none|
|amount|string(int64)|false|none|none|
|isBuyerMaker|boolean|false|none|none|

<h2 id="tocS_klineAggtradeSannerKLineResp">klineAggtradeSannerKLineResp</h2>
<!-- backwards compatibility -->
<a id="schemaklineaggtradesannerklineresp"></a>
<a id="schema_klineAggtradeSannerKLineResp"></a>
<a id="tocSklineaggtradesannerklineresp"></a>
<a id="tocsklineaggtradesannerklineresp"></a>

```json
{
  "status": "string",
  "desc": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|status|string|false|none|none|
|desc|string|false|none|none|

<h2 id="tocS_klineGetAggtradeResponse">klineGetAggtradeResponse</h2>
<!-- backwards compatibility -->
<a id="schemaklinegetaggtraderesponse"></a>
<a id="schema_klineGetAggtradeResponse"></a>
<a id="tocSklinegetaggtraderesponse"></a>
<a id="tocsklinegetaggtraderesponse"></a>

```json
{
  "desc": "string",
  "status": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|desc|string|false|none|none|
|status|string|false|none|none|

<h2 id="tocS_klineGetKlineResponse">klineGetKlineResponse</h2>
<!-- backwards compatibility -->
<a id="schemaklinegetklineresponse"></a>
<a id="schema_klineGetKlineResponse"></a>
<a id="tocSklinegetklineresponse"></a>
<a id="tocsklinegetklineresponse"></a>

```json
{
  "kline": [
    {
      "open_time": 0,
      "open": "string",
      "high": "string",
      "low": "string",
      "close": "string",
      "volume": "string",
      "close_time": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|kline|[[GetKlineResponseKline](#schemagetklineresponsekline)]|false|none|none|

<h2 id="tocS_klineGetTradeRecentResponse">klineGetTradeRecentResponse</h2>
<!-- backwards compatibility -->
<a id="schemaklinegettraderecentresponse"></a>
<a id="schema_klineGetTradeRecentResponse"></a>
<a id="tocSklinegettraderecentresponse"></a>
<a id="tocsklinegettraderecentresponse"></a>

```json
{
  "trade": [
    {
      "id": "string",
      "time": "string",
      "price": "string",
      "amount": "string",
      "isBuyerMaker": true
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|trade|[[GetTradeRecentResponseTrade](#schemagettraderecentresponsetrade)]|false|none|none|

<h2 id="tocS_klineStopPullAggtradeResponse">klineStopPullAggtradeResponse</h2>
<!-- backwards compatibility -->
<a id="schemaklinestoppullaggtraderesponse"></a>
<a id="schema_klineStopPullAggtradeResponse"></a>
<a id="tocSklinestoppullaggtraderesponse"></a>
<a id="tocsklinestoppullaggtraderesponse"></a>

```json
{
  "status": "string",
  "desc": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|status|string|false|none|none|
|desc|string|false|none|none|

<h2 id="tocS_protobufAny">protobufAny</h2>
<!-- backwards compatibility -->
<a id="schemaprotobufany"></a>
<a id="schema_protobufAny"></a>
<a id="tocSprotobufany"></a>
<a id="tocsprotobufany"></a>

```json
{
  "type_url": "string",
  "value": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|type_url|string|false|none|none|
|value|string(byte)|false|none|none|

<h2 id="tocS_rpcStatus">rpcStatus</h2>
<!-- backwards compatibility -->
<a id="schemarpcstatus"></a>
<a id="schema_rpcStatus"></a>
<a id="tocSrpcstatus"></a>
<a id="tocsrpcstatus"></a>

```json
{
  "code": 0,
  "message": "string",
  "details": [
    {
      "type_url": "string",
      "value": "string"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int32)|false|none|none|
|message|string|false|none|none|
|details|[[protobufAny](#schemaprotobufany)]|false|none|none|

