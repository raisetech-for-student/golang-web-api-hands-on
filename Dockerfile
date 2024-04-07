FROM golang:1.22
ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app
COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go install mvdan.cc/gofumpt@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
