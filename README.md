### Run

To run services just execute next command: `$ docker-compose up`

### Run Tests
Due to the lack of time, project has only one test :,(((

To run test: `go test github.com/walkline/shippingpg/clientapi/scanner/json`

### Requests sample

To import port data:
```
curl --location --request POST 'http://localhost:8080/v1/ports/import' \
--header 'Content-Type: application/json' \
--data-binary '@ports.json'
```
To find port by id:
```
curl 'http://localhost:8080/v1/ports?id=BRARB'
```
