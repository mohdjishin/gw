FROM golang:latest AS builder


WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o us

FROM alpine:latest

COPY --from=builder /app/us /app/public.pem ./

EXPOSE 8082

CMD [ "./us" ]
