FROM golang:1.21-bullseye as build

WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata && apt-get clean
ENV TZ="America/Sao_paulo"

CMD ["/myapp"]