FROM golang:1.19-alpine3.16

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN go build -o /daily-kural

CMD ["/daily-kural"]