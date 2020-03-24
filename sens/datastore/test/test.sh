# auth
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/create", "Method": "POST", "Body": {"Email": "emayank@gmail.com", "Mobile": "+917032806003", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/create", "Method": "POST", "Body": {"AuthId": "3c6f2898-a1e6-453d-82a7-dafc2b21f8c7", "Name": "mindfit", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# user
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/users/create", "Method": "POST", "Body": {"AuthId": "3c6f2898-a1e6-453d-82a7-dafc2b21f8c7", "Name": "mindfit", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/users/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org_users
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-users/create", "Method": "POST", "Body": {"OrgId": "9daf33f2-da5e-4b6d-a45b-ec8199b66620", "UserId": "853b1233-12dc-4e02-8a34-6be65994db79", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-users/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org_ops
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-ops/create", "Method": "POST", "Body": {"OrgId": "9f87818f-e3d8-4afe-a1af-deecb87dc4ea", "UserId": "c5d56fa2-fcb5-4dc0-ade1-c0fbb0ac288c", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-users/find?limit=100", "Method":"GET"}' http://35.238.76.35/test/ds
