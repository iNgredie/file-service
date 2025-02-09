# Modules Caching
FROM golang:1.23 as modules

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Build
FROM golang:1.23

WORKDIR /app

COPY --from=modules /go/pkg /go/pkg

COPY . .

RUN CGO_ENABLED=0 go build -o /my-app ./cmd/app

ENTRYPOINT ["/my-app"]
