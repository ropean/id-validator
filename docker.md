```bash
# 构建镜像
make docker-build

# 启动
make docker-up

# 浏览器打开
http://localhost:8080
```

验证容器是否以非 root 运行：

```bash
docker compose exec id-validator whoami
# 应输出: appuser
```

看日志确认启动正常：

```bash
make docker-logs
# 应输出: Server listening on http://localhost:8080
```

停掉：

```bash
make docker-down
```