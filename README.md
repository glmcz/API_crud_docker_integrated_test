
### The simplest http server
- Server with default templates can be started on port :3000 by this cmd:
  - make run-local
- The app waits for Postgres to start
- handle post and get req
- GitHub action included

## How to run inside Docker?
- make build-linux
- make compose-up
- Or inside IDE GoLand.

## Tested with JSON below:
```
    {
      "threatName": "Win32/Rbot",
      "category": "trojan",
      "size": 437289,
      "detectionDate": "2019-04-01",
      "variants": [
        {
          "name": "Win32/TrojanProxy.Emotet.A",
          "dateAdded": "2019-04-10"
        },
        {
          "name": "Win32/TrojanProxy.Emotet.B",
          "dateAdded": "2019-04-22"
        }
      ]
    }
```


## Tested with curl below:

curl -X POST -H "Content-Type: application/json" \                                                                                                         ⬢  system at 12:44:44
-d '{ "id": "224e9a8e-5571-48d3-9da4-c18a1974e268",
"name": "John Doe",
"email": "john.doe@example.com",
"date_of_birth": "1990-01-01"
}' http://localhost:3000/save


curl -X GET "http://localhost:3000/224e9a8e-5571-48d3-9da4-c18a1974e268"
