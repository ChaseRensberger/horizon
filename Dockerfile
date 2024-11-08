FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

# COPY .env.* ./

COPY core/ ./core/

COPY models/ ./models/

COPY config/ ./config/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD ["/main"]