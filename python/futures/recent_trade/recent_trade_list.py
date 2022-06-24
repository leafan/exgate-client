import requests

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/recent_trades?symbol=btcusdt&limit=10"

payload={}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
