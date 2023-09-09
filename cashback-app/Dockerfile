FROM golang

WORKDIR /app/src/cashback

ENV GOPATH=/app
COPY app /app/src/cashback

RUN go build main.go

ENTRYPOINT ["./main"]

EXPOSE 8080