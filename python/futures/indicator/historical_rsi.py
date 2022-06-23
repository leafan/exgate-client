import requests
domain = "http://exgate.leafan.cc/api/"
url = domain+"fexgate/v1/fbinance/kline/rsi?limit=50&symbol=btcusdt&interval=5m&signal=9&isProxy=true&period=14"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
