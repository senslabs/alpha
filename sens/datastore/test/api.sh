curl -v -H "Content-Type: application/json" -d '{"Name": "SENS-PXY"}' http://api.senslabs.io/api/devices/create?state=org
curl -v -H "Content-Type: application/json" -d '{"OrgId": "994b0739-fe71-4377-a265-10cb89c95c5e"}' http://api.senslabs.io/api/devices/a0e1e1de-0e3e-4057-b60a-9173b6a69868/register?state=user

curl -v -H "Content-Type: application/json" -d '{"Name": "SENS-PXY"}' http://api.senslabs.io/api/devices/create?state=org