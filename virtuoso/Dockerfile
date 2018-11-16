FROM golang:latest as builder
COPY healthcheck.go .

RUN go build healthcheck.go

FROM tenforce/virtuoso:virtuoso7.2.5
COPY --from=builder /go/healthcheck /

HEALTHCHECK --start-period=60s CMD ["/healthcheck"]
