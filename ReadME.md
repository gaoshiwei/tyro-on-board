# 创建Person
curl -X POST http://localhost:9000/insert/person -H "ContentType:application/json" -d '{"username":"test1", "sex":"man", "email":"test1@qq.com"}'
控制台日志 insert success: number，number 为自增user_id int, 后续接口可以使用number值作为user_id

# 查询Person
curl http://localhost:9000/select/person?user_id=5

# 更新Person username
curl -X POST http://localhost:9000/update/person -H "ContentType:application/json" -d '{"userid":5 ,"username":"test11", "sex":"man", "email":"test1@qq.com"}'

# 删除与查询、更新类似