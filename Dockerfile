FROM golang:1.22 AS build
ENV CGO_ENABLED=0
WORKDIR /redirect-server/
COPY . .
RUN pwd && go build -v . && ls -la

FROM scratch AS run
EXPOSE 8080
COPY --from=build /redirect-server/redirect-server /redirect-server
