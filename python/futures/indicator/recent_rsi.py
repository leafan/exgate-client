import requests

domain = "http://exgate.leafan.cc/api/"
url = domain+"fexgate/v1/fbinance/kline/historical/rsi?symbol=btcusdt&interval=5m&startTime=1655019000000&endTime=1655043300000&signal=9&isProxy=false&period=14"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
