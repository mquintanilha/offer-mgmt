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

url = "http://sandbox-api.lomadee.com/v3/" + sys.argv[1] + "/offer/_search?keyword="+ sys.argv[2]


# Were SYS.ARG[1] is TOKEN APPLICATION and SYS.ARGV[2] is CATEGORY FOR SEARCHE.

response = requests.get(url, headers=headers, params=querystring)
jsonRes = response.json()

with open('data.json', 'a') as outfile:
    json.dump(jsonRes['offers'], outfile, indent=4, separators=(',',':'))

x = json.load(open('data.json', 'r'))

y = {"installment": { "value": "in stock", "quantity": "new" }}

z = 

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
