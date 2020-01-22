import requests
import json
import sys
import csv

querystring = {"sourceId":"22830830"}
headers = {
    'Accept': "*/*",
    'Cache-Control': "no-cache",
    'Host': "sandbox-api.lomadee.com",
    'Accept-Encoding': "gzip, deflate",
    'Connection': "keep-alive",
    'cache-control': "no-cache"
    }

url = "http://sandbox-api.lomadee.com/v3/1578677161277ae2a08a2/offer/_search?keyword="+ sys.argv[1]

response = requests.get(url, headers=headers, params=querystring)
jsonRes = response.json()

with open('data.json', 'a') as outfile:
    json.dump(jsonRes['offers'], outfile, indent=4, separators=(',',':'))

x = json.load(open('data.json', 'r'))

y = {"installment": { "value": "in stock", "quantity": "new" }}

f = csv.writer(open('import.csv', 'wb+'))

f.writerow(['id', 'title', 'description', 'availability', 'condition', 'price', 'link', 'image_link', 'brand', 'additional_image_link', 'google_product_category'])

for x in x:
    f.writerow([x['id'],
              x['name'],
              x['name'],
              y['installment']['value'],
              y['installment']['quantity'],
              x['price'],
              x['link'],
              x['thumbnail'],
              x['store']['name'],
              x['store']['thumbnail'],
              x['id']])
