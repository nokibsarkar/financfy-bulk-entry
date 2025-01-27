FROM golang:1.22-alpine

RUN apk add --no-cache make curl

WORKDIR /usr/src/app
COPY . .

RUN go mod tidy

CMD ["make", "run"]

