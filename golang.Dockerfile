FROM golang:1.19.4-bullseye AS build
WORKDIR /app/src/api-sample
ENV GOPATH=/app
COPY . .
RUN go mod download
RUN go build -o apisample

FROM gcr.io/distroless/base-debian11 AS deploy
WORKDIR /
COPY --from=build /app/src/api-sample/apisample ./
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/apisample"]

FROM deploy AS final