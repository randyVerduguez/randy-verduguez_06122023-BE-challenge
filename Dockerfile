FROM golang:1.21.5

EXPOSE 8080

WORKDIR /src

COPY .env .

COPY . .

RUN go mod download && go mod verify

RUN go build -o main cmd/app/main.go

ENTRYPOINT [ "./main" ]