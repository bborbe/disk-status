# Disk-Status

Provides disk metrics for Prometheus.

## Run

```
docker run \
--publish 8080:8080 \
--env PORT=8080 \
--env PATH=/ \
docker.io/bborbe/disk-status:latest
```
