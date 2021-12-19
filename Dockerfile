FROM golang:1.17.5 as builder
COPY ./ /fullcycle-desafio-go
WORKDIR /fullcycle-desafio-go
RUN go test && go build

FROM scratch
COPY --from=builder /fullcycle-desafio-go/fullcycle-desafio-go /app
CMD ["/app"]