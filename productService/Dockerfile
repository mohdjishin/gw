FROM golang:latest as builder


WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ps

FROM alpine:latest

COPY --from=builder /app/ps /app/public.pem ./

EXPOSE 8082

CMD [ "./ps" ]
