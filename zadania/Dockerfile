FROM golang:1.14.0-alpine

WORKDIR /app

ADD . /app

RUN go mod download

# Command to run the executable
CMD ["sh", "entrypoint.sh"]
