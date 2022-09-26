.PHONY: test postgres_init create_db drop_db link_db redis_init migrate_install migrate_init_db migrate_up migrate_down sqlc format swag pull server_init docker_build docker_run run run_back callvis
test: # 运行所有的测试程序
	 go test -v -cover ./... -count=1
postgres_zr_init: # 初始化postgres数据库
	docker run --name postgres_zr -v ttms_postgres_zr_data:/var/lib/postgresql/data -v 项目路径/config/postgres/my_postgres.conf:/etc/postgresql/postgresql.conf -p 5432:5432 -e ALLOW_IP_RANGE=0.0.0.0/0 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=ttms -d chenxinaz/zhparser -c 'config_file=/etc/postgresql/postgresql.conf'
drop_db: # 删除数据库
	docker exec -it postgres_zr dropdb ttms
link_db:  # 建立数据库连接
	docker exec -it postgres_zr psql -U root ttms
redis_init: # redis初始化
	docker run --name redis_62 --privileged=true -p 7963:7963 -v ttms_redis_data:/redis/data -v 项目路径/ttms/config/redis:/etc/redis -d redis:6.2 redis-server /etc/redis/redis.conf
redis_link: # redis链接
	docker exec -it redis_62 redis-cli -p 7963
migrate_install: # 安装migrate
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz & sudo mv migrate /usr/bin/migrate
migrate_init_db: # 初始化数据库
	migrate create -ext sql -dir internal/dao/db/migration -seq init_schema
migrate_up: # 向上迁移数据库
	migrate -path internal/dao/db/migration -database "postgresql://root:123456@localhost:5432/ttms?sslmode=disable" -verbose up
migrate_up1:
	migrate -path internal/dao/db/migration -database "postgresql://root:123456@localhost:5432/ttms?sslmode=disable" -verbose up 1
migrate_down:
	migrate -path internal/dao/db/migration -database "postgresql://root:123456@localhost:5432/ttms?sslmode=disable" -verbose down
migrate_down1:
	migrate -path internal/dao/db/migration -database "postgresql://root:123456@localhost:5432/ttms?sslmode=disable" -verbose down 1
migrate_version:
	migrate -path internal/dao/db/migration -database "postgresql://root:123456@localhost:5432/ttms?sslmode=disable" version
sqlc: # sqlc生成go代码
	sqlc generate
goimports_install: # goimports安装
	go get golang.org/x/tools/cmd/goimports
format: # 格式化并检查代码
	goimports -w . && gofmt -w . && golangci-lint run
install_golang-cli: # 安装golang-cli工具，用于静态检查代码质量
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
swag_install: # 安装swag工具，用于生成swagger文档
	go install github.com/swaggo/swag/cmd/swag@v1.8.0
swag: # swag生成文档
	swag init && swag fmt
pull: # 拉取并变基代码
	git fetch origin master && git rebase origin/master
server_init: # 初始化server
	docker start postgres_zr redis_62
docker_build: # 构建docker镜像
	docker build -t ttms:app .
docker_run: # docker运行镜像
	docker run -d --name ttms_app --network ttms_net -p 8080:8080 ttms:app
docker_net: # 创建docker网络
	docker network create ttms_net
run: # 运行server
	go build -o bin/main main.go && ./bin/main
run_back: # 后台运行
	go build -o bin/main main.go && nohup ./bin/main > nohup.out &
callvis: # 调用图
	go-callvis -file=docs/dial -ignore github.com/gin-gonic/gin -format=svg -nostd -nointer -include . -group pkg,type ttms
