# fetch-webpage
This is example fetch url via command line

### How to run application:
- Local Development
    - install golang with version `1.22`
    - run `go mod tidy` for install all related library
    - if using `go run`, 
        - run `go run cmd/main.go fetch {url} {url} ...` to fetch multiple page 
        - run `go run cmd/main.go fetch --metadata url` for url metadata
    - if using `go build`, run`go build -o bin/app cmd/main.go`
        - run `./bin/app fetch {url} {url} ...` to fetch multiple page
        - run `./bin/app fetch --metadata url` for url metadata
    - you can check html files in folder `files/html/*`
    - you can check metadata files in folder `files/json/fetch_data.json`
    - for testing you can run `go test ./... -v`

- Docker
    - run command `docker build -t app .`
    - run command `docker volume create app_volume` to create volume
    - run command `docker run --rm -v app_volume:/usr/bin/files app fetch {url} {url}...` to fetch multiple page
    - run command `docker run --rm -v app_volume:/usr/bin/files app fetch --metadata {url}` for url metadata
    - you can check html and json files in `app_volume` in your docker desktop see volumes section