FROM golang:1.21
ENV TZ=Asia/Tokyo

WORKDIR /app/src

COPY /src /app/src

RUN go mod tidy

CMD ["go","run","cmd/main.go"]