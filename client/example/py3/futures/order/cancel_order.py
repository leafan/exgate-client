import requests
import json

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/cancel?origClientOrderId=0Z1M2GW3lw0UOHMrlOlAXg&symbol=BTCUSDT"

apiKey = ""
apiSecret = ""
payload = {}
headers = {
    'x-exgate-key':
    apiKey,
    'x-exgate-secret':
    apiSecret,
    'Content-Type': 'application/json'
}

response = requests.request("DELETE", url, headers=headers, data=payload)

print(response.text)
