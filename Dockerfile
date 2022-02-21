FROM golang:1.17-alpine

WORKDIR /app

ENV RPC_URL $RPC_URL \
    MAX_BLOCK_DIFFERENCE $MAX_BLOCK_DIFFERENCE

WORKDIR /go/src/app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

ENTRYPOINT ["./main"]
