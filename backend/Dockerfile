## Build metrics backend
FROM golang:1.19 AS builder

WORKDIR /usr/src/app

COPY . .

RUN ls -l

RUN go build -o metrics cmd/metrics/main.go

## Prepare running env
FROM golang:1.19 AS runner

WORKDIR /home/metrics

COPY --from=builder /usr/src/app/metrics /home/metrics

CMD ["./metrics"]
