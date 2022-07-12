FROM golang:1.16-alpine as builder
RUN go version
RUN apk update && apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o bin/go-getting-started -v .

FROM alpine:3
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /src/bin /bin
USER nobody:nobody
CMD [ "/bin/go-getting-started" ]
