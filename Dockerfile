# syntax=docker/dockerfile:1
FROM golang:1.22 AS build-stage

WORKDIR /build

COPY . .

RUN go mod tidy 

RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o app cmd/main.go


FROM alpine:latest

RUN apk update && apk add bash && apk --no-cache add tzdata

WORKDIR /usr/bin

COPY --from=build-stage /build/app .
COPY --from=build-stage /build/files/html files/html
COPY --from=build-stage /build/files/json files/json

RUN chmod +x ./app \
    && chmod -R 777 files
    
ENTRYPOINT ["./app"] --v