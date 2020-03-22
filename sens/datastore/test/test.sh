#auth
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/create", "Method": "POST", "Body": {"Email": "emayank@gmail.com", "Mobile": "+917032806003", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

#org
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/create", "Method": "POST", "Body": {"AuthId": "cf554a64-5557-4bcd-ab20-2ff179e01926", "Name": "mindfit", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# 6f0254dd-3245-4278-ba11-513038ccd9a2