FROM golang:1.16-alpine as builder

RUN apk add git

# Add Maintainer Info
LABEL maintainer="<>"

RUN mkdir /app
ADD . /app
WORKDIR /app/

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download


# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 10000
EXPOSE 8080


# Run Executable
CMD ["./main"]