FROM golang as builder

WORKDIR /app/sentinel

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

LABEL author="Maina Wycliffe"

WORKDIR /root/
COPY --from=builder /go/src/github.com/coding-latte/golang-docker-multistage-build-demo/app .

CMD [ "./app" ]



ENTRYPOINT ["/bin/sentinel"]

