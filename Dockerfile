FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app ./app/main.go


# ------------------------------------


FROM alpine:latest

WORKDIR /kemasan

COPY --from=builder /app/main.app .

CMD [ "/kemasan/main.app" ]