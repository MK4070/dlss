FROM golang:1.24-alpine3.21 AS build
WORKDIR /go/src/dlss
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/dlss ./cmd/dlss


FROM alpine/socat:latest
COPY --from=build /go/bin/dlss /bin/dlss
ENTRYPOINT ["/bin/dlss"]