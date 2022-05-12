FROM golang:1.18.1-bullseye as builder

WORKDIR /CourseWork

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./main.go

FROM alpine:3
COPY --from=builder main /bin/main
COPY --from=builder CourseWork/logger/logs logger/logs
COPY --from=builder CourseWork/gen/upload gen/upload

EXPOSE 30777
ENTRYPOINT ["/bin/main"]