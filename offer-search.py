import requests
import json
import sys
import csv

querystring = {"sourceId":"<ID_AFILIADO>"}

host_sandbox = 'sandbox-api.lomadee.com'
host_prod = 'api.lomadee.com'

if sys.argv[3] == 'prod':
    url = "http://" + host_prod + "/v3/" + sys.argv[1] + "/offer/_search?keyword=" + sys.argv[2]
    headers = {
        'Accept': "*/*",
        'Cache-Control': "no-cache",
        'Host': host_prod,
        'Accept-Encoding': "gzip, deflate",
        'Connection': "keep-alive",
        'cache-control': "no-cache"
        }
elif sys.argv[3] == 'dev':
    url = "http://" + host_sandbox + "/v3/" + sys.argv[1] + "/offer/_search?keyword=" + sys.argv[2]
    headers = {
        'Accept': "*/*",
        'Cache-Control': "no-cache",
        'Host': host_sandbox,
        'Accept-Encoding': "gzip, deflate",
        'Connection': "keep-alive",
        'cache-control': "no-cache"
        }

# Were SYS.ARG[1] is TOKEN APPLICATION and SYS.ARGV[2] is CATEGORY FOR SEARCH and SYS.ARGV[3] is enviroment code in this case only prod or dev.

response = requests.get(url, headers=headers, params=querystring)
jsonRes = response.json()

with open('data.json', 'a') as outfile:
    json.dump(jsonRes['offers'], outfile, indent=4, separators=(',',':'))

x = json.load(open('data.json', 'r'))

y = {"installment": { "value": "in stock", "quantity": "new" }}

z = {"g_id": 772}

f = csv.writer(open('import.csv', 'wb+'))

f.writerow(['id', 
            'title', 
            'description', 
            'availability', 
            'condition', 
            'price', 
            'link', 
            'image_link', 
            'brand', 
            'additional_image_link', 
            'google_product_category'])

for x in x:
    f.writerow([x['id'],
                x['name'],
                x['name'],
                y['installment']['value'],
                y['installment']['quantity'],
                x['price'],
                x['store']['link'],
                x['store']['thumbnail'],
                x['store']['name'],
                x['store']['thumbnail'],
                z['g_id']])