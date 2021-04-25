FROM golang:1.15-alpine as dev

WORKDIR /workspace

FROM golang:1.15-alpine as build

WORKDIR /app
COPY * /app/
RUN go build -o app


FROM alpine as runtime
COPY --from=build /app/app /usr/local/bin/app
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]
