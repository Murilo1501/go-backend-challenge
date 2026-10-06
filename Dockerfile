
FROM golang:1.27

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server ./cmd/
EXPOSE 8080
CMD ["./cmd"]