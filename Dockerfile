FROM golang:1.8

WORKDIR /go/src/github.com/docker-monitor/consulgw

COPY . .

RUN go get -d -v ./...

# RUN go install

CMD ["consulgw"]
