start etcd
start redis-server
start nats-server
start jaeger-all-in-one --collector.zipkin.http-port=9411
start micro --registry=etcd --registry_address=localhost:2379 web