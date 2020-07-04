FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o printer .
CMD ["./printer", "--no-print", "--no-delete"]
