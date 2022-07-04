import requests
domain = "http://exgate.leafan.cc/api/"
url = domain+"v1/fbinance/aggtrade/stop?symbol=btcusdt&interval=5m"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
