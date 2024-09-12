# IP 地址归属地查询

基于 ip2region 的 api 实现

IP2Region [https://github.com/lionsoul2014/ip2region](https://github.com/lionsoul2014/ip2region)

### 使用 docker 部署

```
docker run -d --name ip2region -p 8080:8080 moys3389/ip2region-api:latest

# 添加允许跨域访问域名
docker run -d --name ip2region -e CORS=http://foo.com,http://bar.com -p 8080:8080 moys3389/ip2region-api:latest

# 允许全部域名跨域访问
docker run -d --name ip2region -e CORS=* -p 8080:8080 moys3389/ip2region-api:latest
```

### 使用 API

```
# 查看版本
curl 127.0.0.1:8080/api/version

# 查询自己IP
curl 127.0.0.1:8080/api/search

# 指定IP查询(GET)
curl 127.0.0.1:8080/api/search?ip=1.1.1.1

# 指定IP查询(POST)
curl -X POST -H "Content-Type: application/json" -d '{"ip":"1.1.1.1"}' 127.0.0.1:8080/api/search

# 批量IP查询(GET)
curl 127.0.0.1:8080/api/batch-search?ip=1.1.1.1,2.2.2.2,3.3.3.3

# 批量IP查询(POST)
curl -X POST -H "Content-Type: application/json" -d '{"ips":["1.1.1.1","2.2.2.2","3.3.3.3"]}' 127.0.0.1:8080/api/batch-search
```
