# build ===========================
FROM golang:1.18.10-bullseye as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# dev ===========================
FROM golang:1.18.10 as dev
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
