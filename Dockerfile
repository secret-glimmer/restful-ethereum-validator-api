FROM golang:1.23.0

# Set destination for COPY
WORKDIR /opt/app

# Prepare swag cli
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN export PATH=$(go env GOPATH)/bin:$PATH

# Copy the source code.
COPY . .

RUN swag init -g ./cmd/main.go

# Download Go modules
RUN go mod download

RUN go build -o app ./cmd/...

ENTRYPOINT [ "/opt/app/app" ]