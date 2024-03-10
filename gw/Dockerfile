FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o gw

FROM alpine:latest

COPY --from=builder /app/gw /app/config.yml /app/ec-priv.pem ./
EXPOSE 8080
CMD [ "./gw" ]
