############################
# STEP 1 build executable binary
############################
FROM golang:1.15.8-alpine3.13 AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV SOURCES /go/src/github.com/puxin71/advanced-cloudnative-go/Frameworks/Gin-Web
COPY . ${SOURCES}
WORKDIR ${SOURCES}

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
-ldflags='-w -s -extldflags "-static"' -a \
-o /go/bin/Gin-Web .

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/Gin-Web /go/bin/Gin-Web

# Run the hello binary.
ENTRYPOINT ["/go/bin/Gin-Web"]