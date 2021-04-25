# go-cli-demo

A simple CLI application written in Golang

## Build

```bash
docker build . -t network-cli-demo

```

## Run

```bash
docker run -it network-cli-demo --help
docker run -it network-cli-demo ns google.com
docker run -it network-cli-demo ip google.com
docker run -it network-cli-demo mx google.com
```
