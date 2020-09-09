FROM golang:alpine

RUN mkdir /app
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
WORKDIR /app/src
RUN go build -o ../main
WORKDIR /app
ENTRYPOINT ["./main"]
