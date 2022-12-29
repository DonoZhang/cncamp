FROM golang AS builder
WORKDIR $GOPATH/src/httpserver/
COPY main.go ./main.go
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/httpserver 

FROM scratch
ARG version
WORKDIR /bin/
COPY --from=builder /bin/httpserver ./httpserver
ENV VERSION=$version
EXPOSE 8080
ENTRYPOINT ["./httpserver"]