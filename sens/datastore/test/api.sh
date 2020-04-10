curl -v -H "Content-Type: application/json" -d '{"Name": "SENS-PXY"}' http://api.senslabs.io/api/devices/create?state=org
curl -v -H "Content-Type: application/json" -d '{"OrgId": "994b0739-fe71-4377-a265-10cb89c95c5e"}' http://api.senslabs.io/api/devices/a0e1e1de-0e3e-4057-b60a-9173b6a69868/register?state=user

curl -v -H "Content-Type: application/json" -d '{"Name": "SENS-PXY"}' http://api.senslabs.io/api/devices/create?state=org

curl -kv -d '{"Medium" : "Mobile", "MediumValue" : "+919969129910"}' -H "Content-Type: application/json" -H "x-sens-api-key-id:cf812825-b2c8-48b7-b934-1022ef3c4293" -H "x-sens-api-key:71981e558c90ef07ba4364b8c0174b87" -X POST "https://ws.senslabs.io/api/otp/request"

curl -kv -d '{"Email": "emayank@gmail.com", "Mobile": "+917032806003", "Social": "FB", "FirstName": "Mayank", "LastName": "Joshi", "UserId":"8047443e-10d2-41ac-9889-2683e48fecc7"}' http://localhost:9806/api/users/signup