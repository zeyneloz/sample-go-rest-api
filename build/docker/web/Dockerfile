FROM golang:alpine as builder

RUN apk update && apk add git

RUN mkdir /build
ADD . /build/
WORKDIR /build

# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

FROM golang

WORKDIR /app
COPY --from=builder build/main .

# ENTRYPOINT [ "bash", "-c", "'while !</dev/tcp/postgres_main_db/5432; do sleep 1; done; ./main'" ]
ENTRYPOINT [ "./main" ]