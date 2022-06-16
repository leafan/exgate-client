import requests
import json

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/order"

apiKey = ""
apiSecret = ""
payload = json.dumps({
    "symbol": "BTCUSDT",
    "side": "BUY",
    "type": "LIMIT",
    "price": "10000",
    "amount": "0.001",
    "options": {
        "positionSide": "LONG",
        "workingType": "CONTRACT_PRICE"
    }
})

headers = {
    'x-exgate-key':
    apiKey,
    'x-exgate-secret':
    apiSecret,
    'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
