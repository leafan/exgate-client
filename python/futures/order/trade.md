---
title: api/exgate/v1/futures/trade/trade.proto version not set
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

## FutureTrade_FutureCancelOrder

<a id="opIdFutureTrade_FutureCancelOrder"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.delete('/fexgate/v1/{exchange}/cancel', headers = headers)

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
    req, err := http.NewRequest("DELETE", "/fexgate/v1/{exchange}/cancel", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /fexgate/v1/{exchange}/cancel`

<h3 id="futuretrade_futurecancelorder-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|query|string|false|required 订单ID.|
|orderId|query|string|false|required 订单ID.|
|origClientOrderId|query|string|false|用户订单ID.|

> Example responses

> 200 Response

```json
{
  "cumqty": "string",
  "origqty": "string",
  "side": "string",
  "status": "string",
  "price": "string"
}
```

<h3 id="futuretrade_futurecancelorder-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[tradeCancelOrderResponse](#schematradecancelorderresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

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

## FutureTrade_FuturePlaceOrder

<a id="opIdFutureTrade_FuturePlaceOrder"></a>

> Code samples

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/fexgate/v1/{exchange}/order', headers = headers)

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
    req, err := http.NewRequest("POST", "/fexgate/v1/{exchange}/order", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /fexgate/v1/{exchange}/order`

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

<h3 id="futuretrade_futureplaceorder-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
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
  "clientOrderId": "string",
  "orderId": "string",
  "avgPrice": "string",
  "stopPrice": "string",
  "priceRate": "string",
  "priceProtect": true,
  "status": "string"
}
```

<h3 id="futuretrade_futureplaceorder-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A successful response.|[tradePlaceOrderResponse](#schematradeplaceorderresponse)|
|default|Default|An unexpected error response.|[rpcStatus](#schemarpcstatus)|

<aside class="success">
This operation does not require authentication
</aside>

## FutureTrade_FutureRecentTradesList

<a id="opIdFutureTrade_FutureRecentTradesList"></a>

> Code samples

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/fexgate/v1/{exchange}/recent_trades/{symbol}', headers = headers)

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
    req, err := http.NewRequest("GET", "/fexgate/v1/{exchange}/recent_trades/{symbol}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fexgate/v1/{exchange}/recent_trades/{symbol}`

<h3 id="futuretrade_futurerecenttradeslist-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|exchange|path|string|true|required 交易所名称|
|symbol|path|string|true|required 币种|
|limit|query|string(int64)|false|数据条数.|
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

# Schemas

<h2 id="tocS_BalanceResponseBalance">BalanceResponseBalance</h2>
<!-- backwards compatibility -->
<a id="schemabalanceresponsebalance"></a>
<a id="schema_BalanceResponseBalance"></a>
<a id="tocSbalanceresponsebalance"></a>
<a id="tocsbalanceresponsebalance"></a>

```json
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

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|accountAlias|string|false|none|none|
|asset|string|false|none|none|
|balance|string|false|none|none|
|crossWalletBalance|string|false|none|none|
|crossUnPnl|string|false|none|none|
|availableBalance|string|false|none|none|
|maxWithdrawAmount|string|false|none|none|
|marginAvailable|boolean|false|none|none|
|updateTime|number(double)|false|none|none|

<h2 id="tocS_DepthResponseBidList">DepthResponseBidList</h2>
<!-- backwards compatibility -->
<a id="schemadepthresponsebidlist"></a>
<a id="schema_DepthResponseBidList"></a>
<a id="tocSdepthresponsebidlist"></a>
<a id="tocsdepthresponsebidlist"></a>

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
|bid|[[DepthResponsebills](#schemadepthresponsebills)]|false|none|none|

<h2 id="tocS_DepthResponsebills">DepthResponsebills</h2>
<!-- backwards compatibility -->
<a id="schemadepthresponsebills"></a>
<a id="schema_DepthResponsebills"></a>
<a id="tocSdepthresponsebills"></a>
<a id="tocsdepthresponsebills"></a>

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

<h2 id="tocS_RecentTradesListResponseRecentTradesList">RecentTradesListResponseRecentTradesList</h2>
<!-- backwards compatibility -->
<a id="schemarecenttradeslistresponserecenttradeslist"></a>
<a id="schema_RecentTradesListResponseRecentTradesList"></a>
<a id="tocSrecenttradeslistresponserecenttradeslist"></a>
<a id="tocsrecenttradeslistresponserecenttradeslist"></a>

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

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|balances|[[BalanceResponseBalance](#schemabalanceresponsebalance)]|false|none|none|

<h2 id="tocS_tradeCancelOrderResponse">tradeCancelOrderResponse</h2>
<!-- backwards compatibility -->
<a id="schematradecancelorderresponse"></a>
<a id="schema_tradeCancelOrderResponse"></a>
<a id="tocStradecancelorderresponse"></a>
<a id="tocstradecancelorderresponse"></a>

```json
{
  "cumqty": "string",
  "origqty": "string",
  "side": "string",
  "status": "string",
  "price": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|cumqty|string|false|none|none|
|origqty|string|false|none|none|
|side|string|false|none|none|
|status|string|false|none|none|
|price|string|false|none|none|

<h2 id="tocS_tradeDepthResponse">tradeDepthResponse</h2>
<!-- backwards compatibility -->
<a id="schematradedepthresponse"></a>
<a id="schema_tradeDepthResponse"></a>
<a id="tocStradedepthresponse"></a>
<a id="tocstradedepthresponse"></a>

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
|bid_list|[[DepthResponseBidList](#schemadepthresponsebidlist)]|false|none|none|
|ask_list|[[DepthResponseBidList](#schemadepthresponsebidlist)]|false|none|none|

<h2 id="tocS_tradePlaceOrderResponse">tradePlaceOrderResponse</h2>
<!-- backwards compatibility -->
<a id="schematradeplaceorderresponse"></a>
<a id="schema_tradePlaceOrderResponse"></a>
<a id="tocStradeplaceorderresponse"></a>
<a id="tocstradeplaceorderresponse"></a>

```json
{
  "clientOrderId": "string",
  "orderId": "string",
  "avgPrice": "string",
  "stopPrice": "string",
  "priceRate": "string",
  "priceProtect": true,
  "status": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|clientOrderId|string|false|none|none|
|orderId|string(int64)|false|none|none|
|avgPrice|string|false|none|none|
|stopPrice|string|false|none|none|
|priceRate|string|false|none|none|
|priceProtect|boolean|false|none|none|
|status|string|false|none|none|

<h2 id="tocS_tradeRecentTradesListResponse">tradeRecentTradesListResponse</h2>
<!-- backwards compatibility -->
<a id="schematraderecenttradeslistresponse"></a>
<a id="schema_tradeRecentTradesListResponse"></a>
<a id="tocStraderecenttradeslistresponse"></a>
<a id="tocstraderecenttradeslistresponse"></a>

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
|trades|[[RecentTradesListResponseRecentTradesList](#schemarecenttradeslistresponserecenttradeslist)]|false|none|none|

