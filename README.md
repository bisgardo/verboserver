# Verboserver

Trivial proxy server that logs requests and responses to stdout, written in Go.

*Disclaimer:*

This project was created for debugging purposes and contains absolutely nothing interesting.

## Build

Docker build example:

```shell
docker build -t verboserver --build-arg=build_image=golang:1.17-buster --build-arg=base_image=debian:buster-slim .
```
