FROM golang:1.18

ENV CGO_ENABLED=0

ADD go.mod /go/src/gin_CRUD_server/go.mod
ADD go.sum /go/src/gin_CRUD_server/go.sum
WORKDIR /go/src/gin_CRUD_server
# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download

ADD . /go/src/gin_CRUD_server/
WORKDIR /go/src/gin_CRUD_server/

RUN apt-get update -y && apt-get install -y ca-certificates
RUN go build -o server .

FROM scratch
EXPOSE 3000

COPY ssl.key /etc/ssl/certs/
COPY ssl.crt /etc/ssl/certs/
COPY --from=0 /go/src/gin_CRUD_server/server .

CMD ["/server"]