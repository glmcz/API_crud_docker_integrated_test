

curl -X POST -H "Content-Type: application/json" \
-d '{
"id": "1",
"name": "John Doe",
"email": "john.doe@example.com",
"date_of_birth": "1990-01-01"
}' http://localhost:3000/save


curl -X GET http://localhost:3000/1