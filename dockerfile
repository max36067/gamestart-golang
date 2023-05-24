FROM golang:latest

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -C ./cmd -o app

ENTRYPOINT [ "./cmd/app" ]
