FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .

RUN go build -o main ./cmd/api/transaction/server.go

FROM scratch
COPY --from=build /app/main /main

ENTRYPOINT ["/main"]
