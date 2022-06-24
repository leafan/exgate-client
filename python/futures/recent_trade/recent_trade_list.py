import requests

url = "http://exgate.leafan.cc/api/fexgate/v1/fbinance/recent_trades?symbol=btcusdt&limit=10"

payload={}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
