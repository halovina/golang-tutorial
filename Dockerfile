FROM golang:1.22-alpine

ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8500

CMD [ "./main" ]