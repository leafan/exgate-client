import requests

domain = "http://exgate.leafan.cc/api/"
url = domain+"fexgate/v1/fbinance/aggtrade/scanner/allkline?exchange=fbinance&cron_interval=5m"

payload = {}
headers = {}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
