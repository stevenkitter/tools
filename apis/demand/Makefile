build:
	CGO_ENABLED=0 GOOS=linux go build -o tools .
	docker build -t registry.cn-shanghai.aliyuncs.com/xugopher/tools .
	docker push registry.cn-shanghai.aliyuncs.com/xugopher/tools