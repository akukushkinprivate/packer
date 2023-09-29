# Coding Challenge

## Algorithm
Code of algorithm can be found in the `internal/packer/algorithm.go` file of this repo. 

### Build and run
To build and run locally:
```bash
docker-compose up -d --build
```
Then go to [this page](http://localhost:8000) to interact with API locally via Swagger UI.

To stop:
```bash
docker-compose stop
```

### Test
```bash
go test ./...
```
