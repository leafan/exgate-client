import requests

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/kline/madc?limit=50&symbol=btcusdt&interval=5m&periodFast=16&periodSlow=24&signal=9&isProxy=false"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)