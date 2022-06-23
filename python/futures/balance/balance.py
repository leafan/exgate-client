import requests

url = "http://127.0.0.1:8000/fexgate/v1/fbinance/balance"


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
