import requests
import json

domain = "http://exgate.leafan.cc/api/"
url = domain+"fexgate/v1/fbinance/cancel?origClientOrderId=0Z1M2GW3lw0UOHMrlOlAXg&symbol=BTCUSDT"

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
