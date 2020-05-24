# GO Repo base repo
FROM golang:1.12.0-alpine3.9 as builder

RUN apk add git

# Add Maintainer Info
LABEL maintainer="<>"

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM frolvlad/alpine-gxx:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY /input ./input

# Expose port 5000
EXPOSE 5000

# Run Executable
CMD ["./main"]