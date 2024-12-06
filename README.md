# go-prometheus-grafana

Pull and run prometheus

```sh
docker run -p 9090:9090 -v /path/to/prometheus.yml:/etc/prometheus/prometheus.yml -d prom/prometheus
```

Pull and run Loki

```sh
docker run -d --name=loki -p 3100:3100 -v $(pwd)/loki-config.yml:/etc/loki/local-config.yml grafana/loki:latest -config.file=/etc/loki/local-config.yml
```

Pull and run promtail

```sh
docker run --name=promtail -v /Users/mac/go-prometheus-grafana/logs:/logs -v /Users/mac/go-prometheus-grafana/promtail-config.yml:/etc/promtail/config.yml -d -p 9080:9080  grafana/promtail:latest
```

Pull and run grafana

```sh
docker run --name=grafana -d -p 3001:3000 grafana/grafana
```
