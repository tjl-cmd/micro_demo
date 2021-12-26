protoc --go_out=plugins. --micro_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto
jaeger-client 客户端
Agent 客户端代理
Collector 数据处理
Data Store 数据存储
UI 数据查询与前端展示界面
hystrix-go 熔断