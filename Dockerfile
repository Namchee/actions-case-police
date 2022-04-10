FROM golang:1.17

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -o main .

RUN chmod +x /main

CMD ["/main"]