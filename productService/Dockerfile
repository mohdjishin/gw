FROM golang:latest


WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .


RUN go build -o ps
EXPOSE 8082

CMD [ "./ps" ]
