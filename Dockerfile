# Use Ubuntu 24.04 as base
FROM ubuntu:24.04

# Install dependencies
RUN apt-get update && apt-get install -y \
    golang \
    postgresql-client \
    ca-certificates \
    curl

WORKDIR /app

# Copy Go files and build
COPY . .

RUN go mod tidy && go build -o main .

# Expose port dynamically based on .env
CMD ["./main"]
