

curl -X POST -H "Content-Type: application/json" \                                                                                                         ⬢  system at 12:44:44
-d '{ "id": "224e9a8e-5571-48d3-9da4-c18a1974e268",
"name": "John Doe",
"email": "john.doe@example.com",
"date_of_birth": "1990-01-01"
}' http://localhost:3000/save


curl -X GET "http://localhost:3000/?id=224e9a8e-5571-48d3-9da4-c18a1974e268"
