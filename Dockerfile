FROM golang:1.18.3

WORKDIR /go/src/github.com/hamid/snappfood

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN go test tests/trade_test.go

RUN go build -o /snappfood

EXPOSE 5050

CMD [ "/snappfood" ]