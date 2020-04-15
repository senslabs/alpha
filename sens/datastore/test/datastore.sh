# auth
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/create", "Method": "POST", "Body": {"Email": "emayank@gmail.com", "Mobile": "+917032806003", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/find?limit=1&and=Mobile:%2B917032806003", "Method":"GET"}' http://35.238.76.35/test/ds

curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/create", "Method": "POST", "Body": {"Email": "emayank-op@gmail.com", "Mobile": "+917032806003-op", "Social": "FB-op"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/auths/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/create", "Method": "POST", "Body": {"AuthId": "3c6f2898-a1e6-453d-82a7-dafc2b21f8c7", "Name": "mindfit", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/orgs/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/ops/create", "Method": "POST", "Body": {"AuthId": "d62cb14f-b870-41e4-9df8-a851c5393a4a", "Name": "mindfit-op", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/ops/find?limit=100", "Method":"GET"}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/ops/16803204-9051-4d91-b0c4-ed598a664198/get", "Method":"GET"}' http://35.238.76.35/test/ds

# user
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/users/create", "Method": "POST", "Body": {"AuthId": "3c6f2898-a1e6-453d-82a7-dafc2b21f8c7", "Name": "mindfit", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/users/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org_users
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-users/create", "Method": "POST", "Body": {"OrgId": "994b0739-fe71-4377-a265-10cb89c95c5e", "UserId": "853b1233-12dc-4e02-8a34-6be65994db79", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-users/find?limit=1", "Method":"GET"}' http://35.238.76.35/test/ds

# org_ops
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-ops/create", "Method": "POST", "Body": {"OrgId": "9f87818f-e3d8-4afe-a1af-deecb87dc4ea", "OpId": "16803204-9051-4d91-b0c4-ed598a664198", "Social": "FB"}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/org-ops/find?limit=100", "Method":"GET"}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/endpoints/find?limit=100&and=Category:MQ-TEST-39", "Method":"GET"}' http://35.238.76.35/test/ds

#user_endpoints
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/user-endpoints/create", "Method": "POST", "Body": {"UserId": "853b1233-12dc-4e02-8a34-6be65994db79", "EndpointId": "3db58d1c-3f48-47c7-b5df-a01d88e0dbae", "Access": true}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/user-endpoints/find?limit=100", "Method":"GET"}' http://35.238.76.35/test/ds

#devices
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/devices/create", "Method": "POST", "Body": {"Name": "SENS-RXY", "DeviceId": "28092163-4c3d-4a43-902d-f02a7e03583b", "Status": "CREATED", "Access": true}}' http://35.238.76.35/test/ds
curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/devices/find?limit=100&column=created_at", "Method":"GET"}' http://35.238.76.35/test/ds

curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/devices/11459195-34ce-4ad4-b565-3db59bedfc3c/get", "Method":"GET"}' http://35.238.76.35/test/ds

curl -v -H 'Content-Type: application/json' -d '{"Path":"/api/sessions/create", "Method":"POST", "Body":{"UserId":"10f20136-044e-41ab-b070-fc3d47f023a3", "Name":"PrabhuJiKuchhKaro", "StartedAt":1582569662, "EndedAt":1582569661}}' http://35.238.76.35/test/ds

curl -v -H 'Content-Type: application/json' -d '{"UserId":"10f20136-044e-41ab-b070-fc3d47f023a3", "Name":"test", "StartedAt":1582569662, "EndedAt":1582569661}' http://localhost:9804/api/sessions/create

http://35.238.76.35/test/ds

http://34.71.80.240/api/users/find?limit=2

#------------ datastore APIs
curl -v -H 'Content-Type: application/json' -d '{"EndedAt": 2100}' http://35.184.200.59/api/sessions/92b2c7ff-24ef-480d-a45b-da93d315fed6/update

curl -v -H 'Content-Type: application/json' -d '{"Social": "FBT2"}' http://localhost:9804/api/auths/e4ade338-4c7b-4fd6-8564-5584d1c753bf/update