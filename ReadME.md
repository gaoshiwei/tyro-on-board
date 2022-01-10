# 创建Person
curl -X POST http://localhost:9000/insert/person -H "ContentType:application/json" -d '{"username":"test1", "sex":"man", "email":"test1@qq.com"}'
控制台日志 insert success: number，number 为自增user_id int, 后续接口可以使用number值作为user_id

# 查询Person
curl http://localhost:9000/select/person?user_id=5

# 更新Person username
curl -X POST http://localhost:9000/update/person -H "ContentType:application/json" -d '{"userid":5 ,"username":"test11", "sex":"man", "email":"test1@qq.com"}'

# 删除与查询、更新类似

# 使用docker-compose方式启动服务
- 执行命令docker-compose up 或者 docker-compose up -d (后台执行)
- 执行 curl http://127.0.0.1:9000/， 会返回hello world
- 执行 curl http://127.0.0.1:9000/select/person?user_id=5，使用docker logs -f -t 容器id，可以查看后台日志
- 其他数据库增删改查类推
## 备注
> 为了方便将mysql数据挂在到当前目录的mysql/data下了(- ./mysql/data:/var/lib/mysql)， 自己本地可以改成自己合适的路径