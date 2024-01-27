FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN touch .env

RUN go build -o /todoapi

EXPOSE 8000

CMD ["/todoapi"]

