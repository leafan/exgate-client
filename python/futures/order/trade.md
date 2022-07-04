---
title: api/exgate/v1/spot/trade/exgate.proto version not set
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

<h1 id="api-exgate-v1-spot-trade-exgate-proto">api/exgate/v1/spot/trade/exgate.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="api-exgate-v1-spot-trade-exgate-proto-exgate">ExGate</h1>

## ExGate_SpotBalances

<a id="opIdExGate_SpotBalances"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/balances', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/balances", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/balances`

*获取用户资产*

<h3 id="exgate_spotbalances-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "balances": [
    {
      "asset": "string",
      "available": "string",
      "frozen": "string"
    }
  ]
}
```

<h3 id="exgate_spotbalances-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1BalancesResponse](#schemav1balancesresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_SpotCancelOrder

<a id="opIdExGate_SpotCancelOrder"></a>

> Code samples

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/v1/{exchange}/cancel', headers = headers)

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
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/v1/{exchange}/cancel", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/{exchange}/cancel`

> Body parameter

```json
{
  "symbol": "string",
  "order_id": "string",
  "options": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="exgate_spotcancelorder-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|body|body|object|true|none|
|» symbol|body|string|false|none|
|» order_id|body|string|false|none|
|» options|body|object|false|none|
|»» **additionalProperties**|body|string|false|none|

> Example responses

> 200 Response

```json
{
  "status": "string"
}
```

<h3 id="exgate_spotcancelorder-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[exgatev1CancelOrderResponse](#schemaexgatev1cancelorderresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_SpotDepth

<a id="opIdExGate_SpotDepth"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/depth', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/depth", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/depth`

*获取深度*

<h3 id="exgate_spotdepth-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所|
|symbol|query|string|false|币种.|
|limit|query|string(int64)|false|数量.|
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

<h3 id="exgate_spotdepth-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[exgatev1DepthResponse](#schemaexgatev1depthresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_SpotPlaceOrder

<a id="opIdExGate_SpotPlaceOrder"></a>

> Code samples

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/v1/{exchange}/order', headers = headers)

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
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/v1/{exchange}/order", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/{exchange}/order`

> Body parameter

```json
{
  "symbol": "string",
  "side": "string",
  "type": "string",
  "price": "string",
  "amount": "string",
  "options": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="exgate_spotplaceorder-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|body|body|object|true|none|
|» symbol|body|string|false|none|
|» side|body|string|false|none|
|» type|body|string|false|none|
|» price|body|string|false|none|
|» amount|body|string|false|none|
|» options|body|object|false|may have part of exchange options is difference, so need this options.|
|»» **additionalProperties**|body|string|false|none|

> Example responses

> 200 Response

```json
{
  "order_id": "string"
}
```

<h3 id="exgate_spotplaceorder-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[exgatev1PlaceOrderResponse](#schemaexgatev1placeorderresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_SpotPendingOrders

<a id="opIdExGate_SpotPendingOrders"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/pending_orders', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/pending_orders", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/pending_orders`

<h3 id="exgate_spotpendingorders-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|query|string|false|交易对名称.|

> Example responses

> 200 Response

```json
{
  "orders": [
    {
      "order_id": "string",
      "symbol": "string",
      "type": "string",
      "price": "string",
      "amount": "string",
      "side": "string"
    }
  ]
}
```

<h3 id="exgate_spotpendingorders-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1PendingOrdersResponse](#schemav1pendingordersresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_PullSpotTrade

<a id="opIdExGate_PullSpotTrade"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/pull_trade', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/pull_trade", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/pull_trade`

*拉取现货逐笔交易，走websocket*

<h3 id="exgate_pullspottrade-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|交易所名称|
|symbol|query|string|false|交易对名称.|

> Example responses

> 200 Response

```json
{
  "status": "string",
  "desc": "string"
}
```

<h3 id="exgate_pullspottrade-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[v1PullSpotTradeResponse](#schemav1pullspottraderesponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## ExGate_GetRecentTradesList

<a id="opIdExGate_GetRecentTradesList"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/v1/{exchange}/recent_trades', headers = headers)

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
    req, err := http.NewRequest("GET", "/v1/{exchange}/recent_trades", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{exchange}/recent_trades`

<h3 id="exgate_getrecenttradeslist-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|query|string|false|required 币种.|
|limit|query|string(int64)|false|required 分页页码.|
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

<h3 id="exgate_getrecenttradeslist-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[exgatev1RecentTradesListResponse](#schemaexgatev1recenttradeslistresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_PendingOrdersResponseOrder">PendingOrdersResponseOrder</h2>
<!-- backwards compatibility -->
<a id="schemapendingordersresponseorder"></a>
<a id="schema_PendingOrdersResponseOrder"></a>
<a id="tocSpendingordersresponseorder"></a>
<a id="tocspendingordersresponseorder"></a>

```json
{
  "order_id": "string",
  "symbol": "string",
  "type": "string",
  "price": "string",
  "amount": "string",
  "side": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|order_id|string|false|none|none|
|symbol|string|false|none|none|
|type|string|false|none|none|
|price|string|false|none|none|
|amount|string|false|none|none|
|side|string|false|none|none|

<h2 id="tocS_exgatev1CancelOrderResponse">exgatev1CancelOrderResponse</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1cancelorderresponse"></a>
<a id="schema_exgatev1CancelOrderResponse"></a>
<a id="tocSexgatev1cancelorderresponse"></a>
<a id="tocsexgatev1cancelorderresponse"></a>

```json
{
  "status": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|status|string|false|none|none|

<h2 id="tocS_exgatev1DepthResponse">exgatev1DepthResponse</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1depthresponse"></a>
<a id="schema_exgatev1DepthResponse"></a>
<a id="tocSexgatev1depthresponse"></a>
<a id="tocsexgatev1depthresponse"></a>

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

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|lastUpdateId|number(double)|false|none|none|
|timeStamp|number(double)|false|none|none|
|bid_list|[[exgatev1DepthResponseBidList](#schemaexgatev1depthresponsebidlist)]|false|none|none|
|ask_list|[[exgatev1DepthResponseBidList](#schemaexgatev1depthresponsebidlist)]|false|none|none|

<h2 id="tocS_exgatev1DepthResponseBidList">exgatev1DepthResponseBidList</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1depthresponsebidlist"></a>
<a id="schema_exgatev1DepthResponseBidList"></a>
<a id="tocSexgatev1depthresponsebidlist"></a>
<a id="tocsexgatev1depthresponsebidlist"></a>

```json
{
  "bid": [
    {
      "price": "string",
      "amount": "string"
    }
  ]
}

```

bill_list

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|bid|[[exgatev1DepthResponsebills](#schemaexgatev1depthresponsebills)]|false|none|none|

<h2 id="tocS_exgatev1DepthResponsebills">exgatev1DepthResponsebills</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1depthresponsebills"></a>
<a id="schema_exgatev1DepthResponsebills"></a>
<a id="tocSexgatev1depthresponsebills"></a>
<a id="tocsexgatev1depthresponsebills"></a>

```json
{
  "price": "string",
  "amount": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|price|string|false|none|none|
|amount|string|false|none|none|

<h2 id="tocS_exgatev1PlaceOrderResponse">exgatev1PlaceOrderResponse</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1placeorderresponse"></a>
<a id="schema_exgatev1PlaceOrderResponse"></a>
<a id="tocSexgatev1placeorderresponse"></a>
<a id="tocsexgatev1placeorderresponse"></a>

```json
{
  "order_id": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|order_id|string|false|none|none|

<h2 id="tocS_exgatev1RecentTradesListResponse">exgatev1RecentTradesListResponse</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1recenttradeslistresponse"></a>
<a id="schema_exgatev1RecentTradesListResponse"></a>
<a id="tocSexgatev1recenttradeslistresponse"></a>
<a id="tocsexgatev1recenttradeslistresponse"></a>

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

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|trades|[[exgatev1RecentTradesListResponseRecentTradesList](#schemaexgatev1recenttradeslistresponserecenttradeslist)]|false|none|none|

<h2 id="tocS_exgatev1RecentTradesListResponseRecentTradesList">exgatev1RecentTradesListResponseRecentTradesList</h2>
<!-- backwards compatibility -->
<a id="schemaexgatev1recenttradeslistresponserecenttradeslist"></a>
<a id="schema_exgatev1RecentTradesListResponseRecentTradesList"></a>
<a id="tocSexgatev1recenttradeslistresponserecenttradeslist"></a>
<a id="tocsexgatev1recenttradeslistresponserecenttradeslist"></a>

```json
{
  "id": 0,
  "price": "string",
  "qty": "string",
  "quoteQty": "string",
  "time": 0,
  "isBuyerMaker": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|number(double)|false|none|none|
|price|string|false|none|none|
|qty|string|false|none|none|
|quoteQty|string|false|none|none|
|time|number(double)|false|none|none|
|isBuyerMaker|boolean|false|none|none|

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

<h2 id="tocS_v1BalancesResponse">v1BalancesResponse</h2>
<!-- backwards compatibility -->
<a id="schemav1balancesresponse"></a>
<a id="schema_v1BalancesResponse"></a>
<a id="tocSv1balancesresponse"></a>
<a id="tocsv1balancesresponse"></a>

```json
{
  "balances": [
    {
      "asset": "string",
      "available": "string",
      "frozen": "string"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|balances|[[v1BalancesResponseBalance](#schemav1balancesresponsebalance)]|false|none|none|

<h2 id="tocS_v1BalancesResponseBalance">v1BalancesResponseBalance</h2>
<!-- backwards compatibility -->
<a id="schemav1balancesresponsebalance"></a>
<a id="schema_v1BalancesResponseBalance"></a>
<a id="tocSv1balancesresponsebalance"></a>
<a id="tocsv1balancesresponsebalance"></a>

```json
{
  "asset": "string",
  "available": "string",
  "frozen": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|asset|string|false|none|none|
|available|string|false|none|none|
|frozen|string|false|none|none|

<h2 id="tocS_v1PendingOrdersResponse">v1PendingOrdersResponse</h2>
<!-- backwards compatibility -->
<a id="schemav1pendingordersresponse"></a>
<a id="schema_v1PendingOrdersResponse"></a>
<a id="tocSv1pendingordersresponse"></a>
<a id="tocsv1pendingordersresponse"></a>

```json
{
  "orders": [
    {
      "order_id": "string",
      "symbol": "string",
      "type": "string",
      "price": "string",
      "amount": "string",
      "side": "string"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|orders|[[PendingOrdersResponseOrder](#schemapendingordersresponseorder)]|false|none|none|

<h2 id="tocS_v1PullSpotTradeResponse">v1PullSpotTradeResponse</h2>
<!-- backwards compatibility -->
<a id="schemav1pullspottraderesponse"></a>
<a id="schema_v1PullSpotTradeResponse"></a>
<a id="tocSv1pullspottraderesponse"></a>
<a id="tocsv1pullspottraderesponse"></a>

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

