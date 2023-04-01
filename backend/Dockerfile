## Build metrics backend
FROM golang:1.19 AS builder

WORKDIR /usr/src/app

COPY . .

RUN ls -l

ENV GOPATH /usr/src/app

RUN go build -o metrics backend/cmd/metrics/main.go

## Prepare running env
FROM golang:1.19 AS runner

WORKDIR /home/metrics

COPY --from=builder /usr/src/app/metrics /home/metrics

CMD ["./metrics"]
