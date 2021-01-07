# etcd web page
micro --registry=etcd --registry_address=localhost:2379 web



## web api

set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=localhost:2379
micro api --namespace=go.micro --type=service



## 本地服务端口

### etcd web page

localhost:8082

### jaeger

localhost:16686