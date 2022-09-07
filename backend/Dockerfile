FROM golang:1.19.0-alpine3.16

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make

RUN mkdir /app
ADD . /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run the executable
CMD ["./main"]