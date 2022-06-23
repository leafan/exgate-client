import requests

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/depth?symbol=BTCUSDT&limit=5&isLocal=false&tolerate_interval=1"


payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
