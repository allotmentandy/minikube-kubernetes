FROM golang:1.15-alpine AS base
WORKDIR /src/
COPY hello.go /src/
ENV CGO_ENABLED=0 
RUN go build -o /bin/hello
FROM scratch
COPY --from=base /bin/hello /bin/hello
ENTRYPOINT ["/bin/hello"]