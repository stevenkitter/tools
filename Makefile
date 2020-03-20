all: api work
# 并行执行 make all -j2
# ssh-copy-id root@35.220.159.74 传证书到目标服务器上
api:
	cd ${PWD}/apis/demand; \
	make; \
	echo "demand 接口上传成功"
	ssh root@35.220.159.74 'cd production/api/; \
	docker pull registry.cn-shanghai.aliyuncs.com/xugopher/tools; \
	docker-compose down && docker-compose up -d'
	echo "接口 部署完毕"
work:
	cd ${PWD}/cmd/bank; \
	make; \
	echo "bank 脚本上传成功"
	ssh root@35.220.159.74 'cd production/work/; \
	docker pull registry.cn-shanghai.aliyuncs.com/xugopher/bank; \
	docker-compose down && docker-compose up -d'
	echo "脚本 部署完毕"