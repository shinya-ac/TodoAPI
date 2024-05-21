FROM golang:1.21
ENV TZ=Asia/Tokyo

WORKDIR /app/src

COPY /src /app/src

RUN go mod tidy

WORKDIR /app/src/cmd

ENV GOARCH=amd64

RUN go build -o /build

EXPOSE 8080

CMD ["/build"]