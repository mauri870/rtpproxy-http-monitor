# RTP Proxy Health Check

A simple http server that exposes a /health HTTP endpoint to check the availability of a rtpproxy instance.


## Build

```
go build
```

## Usage

```
./rtpproxyhttpmonitor -addr :8080 -rtpproxy 127.0.0.1:7722
```
