FROM golang:1.21.5

EXPOSE 8080

WORKDIR /src

COPY .env .

COPY . .

# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go mod download && go mod verify

WORKDIR /src/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go run main.go
