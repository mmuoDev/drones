FROM golang:latest

LABEL maintainer="mmuoDev <radioactive.uche11@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum . 

RUN go mod download 

COPY . .

ENV PORT 9064

RUN go build -o main ./cmd/drones

CMD ["./main"]