###################
# BUILD
###################
FROM golang:1.22-alpine As build


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN go build -o main .

###################
# PRODUCTION
###################
FROM alpine:latest AS production

WORKDIR /app

COPY --from=build /app/main .

# ========================================
# This command is for running image on local only
# remove if not using image in local

COPY .env .
RUN export $(cat .env | xargs)
ENV DB_HOST=host.docker.internal
# ========================================

EXPOSE 8080

CMD ["./main"]