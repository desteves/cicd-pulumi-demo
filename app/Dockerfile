# syntax=docker/dockerfile:1
FROM --platform=linux/amd64 golang:1.21.4-alpine3.18  AS build-stage
WORKDIR /app
COPY *.go go.mod go.sum ./
COPY /api /app/api
RUN CGO_ENABLED=0 GOOS=linux go build 

FROM --platform=linux/amd64 golang:1.21.4-alpine3.18 AS build-release-stage
COPY --from=build-stage /app /hello-world-app
ENTRYPOINT ["/hello-world-app"]