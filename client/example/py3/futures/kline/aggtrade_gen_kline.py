import requests

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/aggtrade/scanner/kline?symbol=btcusdt&interval=5m"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
