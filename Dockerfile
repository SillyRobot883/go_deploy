# 1ST STAGE
FROM golang:1.22.5-alpine AS builder

WORKDIR /src
COPY . .

RUN go mod download
RUN go build -o ./test

CMD ["./test"]

# 2ND STAGE: copy first stage into 2nd build to make it lightweight
FROM scratch

WORKDIR /app
COPY --from=builder /src/test ./test

CMD ["/app/test"]