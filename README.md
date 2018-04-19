# Disk-Status

## Metrics

```
disk-status -path /
```

## Docker

```
docker run \
--publish 8080:8080 \
--env PORT=8080 \
--env PATH=/ \
docker.io/bborbe/disk-status:latest \
-logtostderr -v=1
```
