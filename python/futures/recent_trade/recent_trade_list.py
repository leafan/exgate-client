import requests

url = "http://exgate.leafan.cc/api/v1/fbinance/get_recent_trades/btcusdt?limit=20"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
