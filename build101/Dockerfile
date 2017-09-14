FROM golang:1.8.3 as go-builder
ENV CGO_ENABLED=0
ADD main.go .
RUN go build -o /hello

FROM scratch
EXPOSE 80
COPY --from=go-builder /hello /hello
CMD ["/hello"]