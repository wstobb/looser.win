FROM alpine:latest AS builder
RUN apk --no-cache add make go curl bash
RUN curl -fsSL https://bun.com/install | bash
ENV PATH="/root/.bun/bin:$PATH"
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make web

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin .
CMD ["/app/bin"]
