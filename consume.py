import requests
from requests.exceptions import HTTPError

try:
    response = requests.get('http://localhost:8085/path?zipCode=30332&items=Banana&items=Oranges&items=Strawberry')
    response.raise_for_status()
    #access JSON content
    jsonResponse = response.json()
    print("Response", jsonResponse)

except HTTPError as http_err:
    print(f'HTTP error occurred: {http_err}')
except Exception as err:
    print(f'Other error occurred: {err}')