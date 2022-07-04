import requests

url = "http://exgate.leafan.cc/api/v1/fbinance/depth?symbol=BTCUSDT&limit=5&isLocal=false&tolerate_interval=1"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
