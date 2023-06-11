FROM golang:alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /app/cmd/social-media-http ./cmd/social-media-http
RUN apk add && apk add make

EXPOSE 8000

CMD [ "sh", "-c", "/app/cmd/social-media-http/social-media-http"]