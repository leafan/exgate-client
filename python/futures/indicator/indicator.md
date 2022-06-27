---
title: api/exgate/v1/futures/indicator/indicator.proto version not set
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

<h1 id="api-exgate-v1-futures-indicator-indicator-proto">api/exgate/v1/futures/indicator/indicator.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="api-exgate-v1-futures-indicator-indicator-proto-indicators">Indicators</h1>

## Indicators_GetHistoricalMADC

<a id="opIdIndicators_GetHistoricalMADC"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/kline/historical/madc/{symbol}/{interval}/{start_time}/{end_time}/{periodFast}/{periodSlow}/{signal}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/kline/historical/madc/{symbol}/{interval}/{start_time}/{end_time}/{periodFast}/{periodSlow}/{signal}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/kline/historical/madc/{symbol}/{interval}/{start_time}/{end_time}/{periodFast}/{periodSlow}/{signal}`

<h3 id="indicators_gethistoricalmadc-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required，交易所名称|
|symbol|path|string|true|required，交易对名称|
|interval|path|string|true|required，时间周期，可选值：1min, 5min, 15min, 30min, 1hour, 1day|
|start_time|path|string(int64)|true|required，开始时间戳|
|end_time|path|string(int64)|true|required，结束时间戳|
|periodFast|path|string(int64)|true|required，fast信号周期|
|periodSlow|path|string(int64)|true|required，slow信号周期|
|signal|path|string(int64)|true|required，signal信号周期|
|limit|query|string(int64)|false|返回数据条数.|
|isProxy|query|boolean|false|是否使用代理,走http，或者local，默认为false.|

> Example responses

> 200 Response

```json
{
  "macd": [
    {
      "timestamp": "string",
      "madc": 0,
      "signal": 0,
      "histogram": 0
    }
  ]
}
```

<h3 id="indicators_gethistoricalmadc-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[indicatorGetMADCResponse](#schemaindicatorgetmadcresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Indicators_GetHistoricalRSI

<a id="opIdIndicators_GetHistoricalRSI"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/kline/historical/rsi/{symbol}/{interval}/{period}/{start_time}/{end_time}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/kline/historical/rsi/{symbol}/{interval}/{period}/{start_time}/{end_time}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/kline/historical/rsi/{symbol}/{interval}/{period}/{start_time}/{end_time}`

<h3 id="indicators_gethistoricalrsi-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required，交易所名称|
|symbol|path|string|true|required，交易对名称|
|interval|path|string|true|required，时间周期，可选值：1min, 5min, 15min, 30min, 1hour, 1day|
|period|path|string(int64)|true|required，信号周期|
|start_time|path|string(int64)|true|required，开始时间戳|
|end_time|path|string(int64)|true|required，结束时间戳|
|limit|query|string(int64)|false|返回数据条数.|
|isProxy|query|boolean|false|是否使用代理,走http，或者local，默认为false.|

> Example responses

> 200 Response

```json
{
  "rsi": [
    {
      "timestamp": "string",
      "rsi": 0
    }
  ]
}
```

<h3 id="indicators_gethistoricalrsi-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[indicatorGetRSIResponse](#schemaindicatorgetrsiresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Indicators_GetRecentMADC

<a id="opIdIndicators_GetRecentMADC"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/kline/madc/{symbol}/{interval}/{periodFast}/{periodSlow}/{signal}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/kline/madc/{symbol}/{interval}/{periodFast}/{periodSlow}/{signal}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/kline/madc/{symbol}/{interval}/{periodFast}/{periodSlow}/{signal}`

<h3 id="indicators_getrecentmadc-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required，交易所名称|
|symbol|path|string|true|required，交易对名称|
|interval|path|string|true|required，时间周期，可选值：1min, 5min, 15min, 30min, 1hour, 1day|
|periodFast|path|string(int64)|true|required，快速移动平均线周期|
|periodSlow|path|string(int64)|true|required，慢速移动平均线周期|
|signal|path|string(int64)|true|required，信号周期|
|limit|query|string(int64)|false|返回数据条数.|
|isProxy|query|boolean|false|是否使用代理,走http，或者local，默认为false.|

> Example responses

> 200 Response

```json
{
  "macd": [
    {
      "timestamp": "string",
      "madc": 0,
      "signal": 0,
      "histogram": 0
    }
  ]
}
```

<h3 id="indicators_getrecentmadc-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[indicatorGetMADCResponse](#schemaindicatorgetmadcresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## Indicators_GetRecentRSI

<a id="opIdIndicators_GetRecentRSI"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/kline/rsi/{symbol}/{interval}/{period}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/kline/rsi/{symbol}/{interval}/{period}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/kline/rsi/{symbol}/{interval}/{period}`

<h3 id="indicators_getrecentrsi-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required，交易所名称|
|symbol|path|string|true|required，交易对名称|
|interval|path|string|true|required，时间周期，可选值：1min, 5min, 15min, 30min, 1hour, 1day|
|period|path|string(int64)|true|required，信号周期|
|limit|query|string(int64)|false|返回数据条数.|
|isProxy|query|boolean|false|是否使用代理,走http，或者local，默认为false.|

> Example responses

> 200 Response

```json
{
  "rsi": [
    {
      "timestamp": "string",
      "rsi": 0
    }
  ]
}
```

<h3 id="indicators_getrecentrsi-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[indicatorGetRSIResponse](#schemaindicatorgetrsiresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_GetMADCResponseMACD">GetMADCResponseMACD</h2>
<!-- backwards compatibility -->
<a id="schemagetmadcresponsemacd"></a>
<a id="schema_GetMADCResponseMACD"></a>
<a id="tocSgetmadcresponsemacd"></a>
<a id="tocsgetmadcresponsemacd"></a>

```json
{
  "timestamp": "string",
  "madc": 0,
  "signal": 0,
  "histogram": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|timestamp|string(int64)|false|none|none|
|madc|number(double)|false|none|none|
|signal|number(double)|false|none|none|
|histogram|number(double)|false|none|none|

<h2 id="tocS_GetRSIResponseRSI">GetRSIResponseRSI</h2>
<!-- backwards compatibility -->
<a id="schemagetrsiresponsersi"></a>
<a id="schema_GetRSIResponseRSI"></a>
<a id="tocSgetrsiresponsersi"></a>
<a id="tocsgetrsiresponsersi"></a>

```json
{
  "timestamp": "string",
  "rsi": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|timestamp|string(int64)|false|none|none|
|rsi|number(double)|false|none|none|

<h2 id="tocS_indicatorGetMADCResponse">indicatorGetMADCResponse</h2>
<!-- backwards compatibility -->
<a id="schemaindicatorgetmadcresponse"></a>
<a id="schema_indicatorGetMADCResponse"></a>
<a id="tocSindicatorgetmadcresponse"></a>
<a id="tocsindicatorgetmadcresponse"></a>

```json
{
  "macd": [
    {
      "timestamp": "string",
      "madc": 0,
      "signal": 0,
      "histogram": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|macd|[[GetMADCResponseMACD](#schemagetmadcresponsemacd)]|false|none|none|

<h2 id="tocS_indicatorGetRSIResponse">indicatorGetRSIResponse</h2>
<!-- backwards compatibility -->
<a id="schemaindicatorgetrsiresponse"></a>
<a id="schema_indicatorGetRSIResponse"></a>
<a id="tocSindicatorgetrsiresponse"></a>
<a id="tocsindicatorgetrsiresponse"></a>

```json
{
  "rsi": [
    {
      "timestamp": "string",
      "rsi": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|rsi|[[GetRSIResponseRSI](#schemagetrsiresponsersi)]|false|none|none|

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

