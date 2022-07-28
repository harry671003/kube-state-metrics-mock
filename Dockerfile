FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY cmd ./
COPY pkg ./

RUN go mod download

RUN go build -o /kube_state_metrics cmd/server.go 

CMD [ "/kube_state_metrics" ]
