FROM golang:latest
WORKDIR /app
ENV PostgresURI=postgres://postgres:123456@host.docker.internal:5432/userdb?sslmode=disable
ENV PORT=:1000
COPY . .
RUN go mod tidy && \
    go build -o myGolangApp
ENTRYPOINT ./myGolangApp