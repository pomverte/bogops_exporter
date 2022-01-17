# Open metrics playgroung

```
$ go run bogops_exporter.go
$ curl http://127.0.0.1:8080/metrics
```

```
# TYPE ze_bogops_at_octo gauge
# HELP ze_bogops_at_octo Les bogops chez OCTO.
ze_bogops_at_octo{level="Consultant"} 42

# TYPE ze_status_code gauge
# HELP ze_status_code Un status code.
ze_status_code 201
```

## My first golang steps

```
go mod init bogops_exporter

go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp

go mod tidy
```
