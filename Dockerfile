FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Inspect the /go/bin directory (this is where binaries are installed by default)
RUN ls -la /go/bin

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
# Adjust the path to the migrate binary based on the previous inspection
COPY --from=builder /go/bin/migrate /app/migrate
COPY app.env .
COPY start.sh .
COPY db/migrations ./migration
COPY wait-for.sh .
RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
