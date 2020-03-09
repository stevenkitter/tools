build:
	CGO_ENABLED=0 GOOS=linux go build -o bank .
	docker build -t registry.cn-shanghai.aliyuncs.com/xugopher/bank .
	docker push registry.cn-shanghai.aliyuncs.com/xugopher/bank