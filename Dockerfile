from golang:1.23.1

ADD ./src/ /app
WORKDIR /app

RUN go mod download
RUN go build -o api cmd/api/main.go

ENTRYPOINT ./api