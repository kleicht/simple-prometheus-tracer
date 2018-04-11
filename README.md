# simple-prometheus-tracer

```
$ go run main.go
$ hey -n 10000 -c 200 http://localhost:8081/ping
$ curl http://localhost:8081/metrics | grep http_response_time_seconds
```