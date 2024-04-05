# Redis Sentinel Lab

replication info
```
docker compose exec redis-master redis-cli info replication
```

sentinel info
```
docker compose exec sentinel-1 redis-cli -p 26379 info sentinel
```
