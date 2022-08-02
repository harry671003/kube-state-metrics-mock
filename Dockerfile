FROM golang:1.16-alpine

WORKDIR /app

RUN apk add git
ENV GOPROXY direct

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN go build -o /kube_state_metrics ./cmd/server.go 

CMD [ "/kube_state_metrics" ]
