FROM alpine:latest

RUN apk add --no-cache go bash curl jq gcc musl-dev  build-base --repository="http://dl-cdn.alpinelinux.org/alpine/edge/community"

# Arguments
ARG FIREBASE_CONFIG

# CREATE FIREBASE CONFIG
RUN touch firebase.json
RUN echo "${FIREBASE_CONFIG}" >> firebase.json

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
RUN cp /firebase.json .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/main"]