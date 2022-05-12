FROM golang:1.18.1-bullseye as builder

WORKDIR /build

COPY init/go.mod .
COPY init/go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main init/main.go

FROM alpine:3
COPY --from=builder build/init /init
COPY --from=builder main /bin/main
COPY --from=builder build/migrations/sql /migrations/sql
ENTRYPOINT ["/bin/main"]