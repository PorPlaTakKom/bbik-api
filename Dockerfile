FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG DB_USER
ARG DB_PASS
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME

RUN echo "DB_USER=${DB_USER}" > .env && \
    echo "DB_PASS=${DB_PASS}" >> .env && \
    echo "DB_HOST=${DB_HOST}" >> .env && \
    echo "DB_PORT=${DB_PORT}" >> .env && \
    echo "DB_NAME=${DB_NAME}" >> .env

RUN ls -la

RUN cat .env

RUN go build -o main .

EXPOSE 80

CMD ["./main"]
