# 1ST STAGE
FROM golang:1.22.5-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /src
COPY . .

RUN go mod download
ENV CGO_ENABLED=1
RUN go build -o ./app ./cmd/docker_go/main.go || { echo "Go build failed"; exit 1; }

# 2ND STAGE: copy first stage into 2nd build to make it lightweight
FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/app .
COPY users.db .

CMD ["/app/app"]