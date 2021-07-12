FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -tags netgo -o server

FROM scratch
COPY --from=builder /app/server .
EXPOSE 8000
CMD ["./server"]
