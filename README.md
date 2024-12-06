# go-prometheus-grafana

Pull and run prometheus

```sh
docker run -p 9090:9090 -v /path/to/prometheus.yml:/etc/prometheus/prometheus.yml -d prom/prometheus
```

prometheus.yml

```sh
global:
  scrape_interval: 15s
scrape_configs:
  - job_name: 'echo_server'
    static_configs:
      - targets: ['localhost:8080']
```
