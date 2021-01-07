# etcd web page
micro --registry=etcd --registry_address=localhost:2379 web



## web api

set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=localhost:2379
micro api --namespace=go.micro --type=service

