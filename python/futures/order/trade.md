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
|exchange|path|string|true|none|
|sybmol|query|string|false|none|
|orderId|query|string|false|none|
|origClientOrderId|query|string|false|none|

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
|exchange|path|string|true|none|
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

# Schemas

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

