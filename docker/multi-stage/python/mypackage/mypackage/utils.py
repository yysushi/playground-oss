import requests


def get_http_headers():
    res = requests.get('https://httpbin.org/headers')
    return res.json()['headers']
