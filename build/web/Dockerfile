FROM alpine:latest AS builder
RUN apk --no-cache add make go
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make web

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin .
CMD ["/app/bin"]
