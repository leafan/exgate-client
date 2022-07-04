import requests

url = "http://exgate.leafan.cc/api/v1/fbinance/balance"

apiKey = ""
apiSecret = ""
payload = {}
headers = {
    'x-exgate-key':
    apiKey,
    'x-exgate-secret':
    apiSecret
}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
