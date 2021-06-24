FROM golang:1.16 AS builder
WORKDIR /src
COPY . ./
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o server ./cmd

FROM alpine
RUN apk update && \
     apk add libc6-compat && \
     apk add ca-certificates
WORKDIR /

COPY --from=builder /go/src/server .
EXPOSE 9876
CMD ["./server"]
