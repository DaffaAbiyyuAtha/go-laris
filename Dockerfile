FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . /app/

RUN go mod tidy && go build -o /app/backend

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app /app

ENV PATH="/app:${PATH}"

EXPOSE 8080

ENTRYPOINT [ "backend" ]