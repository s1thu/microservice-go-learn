# 1️⃣ Base image
FROM golang:1.24-alpine

# 2️⃣ Set working directory inside container
WORKDIR /app

# 3️⃣ Copy go.mod and go.sum first (cache optimization)
COPY go.mod go.sum ./

# 4️⃣ Download dependencies
RUN go mod download

# 5️⃣ Copy the rest of the source code
COPY . .

# 6️⃣ Build the application
RUN go build -o app ./cmd/main.go

# 7️⃣ Expose API port
EXPOSE 8080

# 8️⃣ Run the binary
CMD ["./app"]