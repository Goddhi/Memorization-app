## implementing multi-stage build
FROM golang:1.19-alpine as builder  

## Development contianer
ENV GO111MODULE=on 

WORKDIR /go/src/app

#  Automatically rebuild and restart the Go application whenever source files change. This is particularly useful for web development, as it allows developers to see changes in real-time without manually restarting the server.
RUN go get github.com/cespare/reflex

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./run .

## second build
## for production i.e this is the binary that will run in the server
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

### copy executable from builder
COPY --from=builder /go/src/app .

EXPOSE 8080
CMD [ "./run" ]



