# ./Dockerfile

FROM golang:1.16-alpine AS builder

# Move to working directory (/build).
RUN apk --no-cache add ca-certificates

WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -installsuffix cgo -o health-check ./healthcheck

# Set necessary environment variables needed for our image 
# and build the API server.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o apiserver .


FROM scratch

# Copy binary and config files from /build 
# to root folder of scratch container.

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/apiserver /
COPY --from=builder /build/health-check /

# Export necessary port.
HEALTHCHECK --interval=10s --timeout=1s --start-period=2s --retries=3 CMD [ "/health-check" ]
EXPOSE 3000

# Necessary environment variables
ENV DB_SERVER_URL=unset MOODLE_HOST=unset SIGNAL_HOST=unset SIGNAL_NUMBER=unset PORT=3000 USE_DOT_ENV=0 CORS=0

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
