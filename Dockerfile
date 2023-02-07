FROM golang:1.19-alpine3.17 as builder

WORKDIR /

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static-debian11
WORKDIR /bot

COPY --from=builder /main /bot/main

USER nonroot:nonroot

ENTRYPOINT ["/bot/main"]