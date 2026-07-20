# 命令参考

## 编译与运行

| 命令 | 说明 |
|---|---|
| `make build` | 编译项目到 `bin/api.exe` |
| `make run` | 编译并运行 |
| `.\scripts\build.ps1` | PowerShell 编译脚本 |
| `bash scripts/build.sh` | Bash 编译脚本 |
| `go build -o bin/api.exe ./cmd/api/` | 原生 Go 编译 |
| `go run ./cmd/api/` | 直接运行 |

## Redis（Docker）

| 命令 | 说明 |
|---|---|
| `.\scripts\redis.ps1` | 启动 Redis（容器存在则启动，不存在则自动创建） |
| `.\scripts\redis.ps1 -Action init` | 删除旧容器，重新创建并启动 |
| `.\scripts\redis.ps1 -Action stop` | 停止 Redis |
| `.\scripts\redis.ps1 -Action restart` | 重启 Redis |
| `docker start redis` | 手动启动已存在的 Redis 容器 |
| `docker run -d --name redis -p 6379:6379 redis` | 手动创建并启动 Redis 容器 |
