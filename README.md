# ova-plan-api

An application to store plans

## How to run

To run application just execute
```make run-with-infra```

For other options see Makefile

## How to use

by default:

* gRPC api - `:8080`. Proto file available [there](/ova-plan-api/blob/main/api/ova-plan-api/api.proto). There is client
  in `/pkg` directory
* prometheus - `:9090`
* jaeger `:16686`