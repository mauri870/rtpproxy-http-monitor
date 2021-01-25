# RTP Proxy Health Check

A simple http server that exposes a /health HTTP endpoint to check the availability of a rtpproxy instance.

This project was intended to be used as a simple health endpoint for private and public status pages that have active http monitoring.


## Build

```
make
make install
```

## Usage

```
./rtpproxy-http-monitor -addr :8080 -rtpproxy 127.0.0.1:7722
```

## Init.d service

An init.d service is available at etc/init.d/rtpproxy-http-monitor.

```bash
make init-d
```
