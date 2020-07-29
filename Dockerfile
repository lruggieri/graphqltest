FROM golang AS build-env

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o myApp .

# final stage
FROM alpine

COPY --from=build-env /app/myApp /
COPY --from=build-env /app/pkg/database /pkg/database

EXPOSE 3000

ENTRYPOINT [ "/myApp" ]