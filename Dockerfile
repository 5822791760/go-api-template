###################
# BUILD
###################
FROM golang:1.22-alpine As build


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o . ./cmd/app

###################
# PRODUCTION
###################
FROM alpine:latest AS production

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 3000

CMD ["./app"]