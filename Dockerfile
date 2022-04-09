FROM golang:1.17

WORKDIR /case-police

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -o main .

RUN chmod +x /case-police/main

CMD ["/case-police/main"]