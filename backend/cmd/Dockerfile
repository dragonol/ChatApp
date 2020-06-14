FROM golang:1.14

WORKDIR /user/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
