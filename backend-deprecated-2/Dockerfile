FROM golang:alpine

WORKDIR /app

COPY . .

# alternative: make a multi-part container. Install goose and apply migrations via cli in the first, go ahead with the rest in the second
RUN go build -o ./bin/ ./cmd/api/ && \
    go build -o ./bin/migrate ./cmd/migrate

CMD ["/app/bin/api"]

