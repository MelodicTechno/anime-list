## 项目结构

anime-list/
├── cmd/
│   └── api/
│       └── main.go          # 应用入口
├── internal/
│   ├── config/              # 配置管理
│   ├── handler/             # HTTP 处理器
│   ├── service/             # 业务逻辑
│   ├── repository/          # 数据访问
│   └── model/               # 数据模型
├── pkg/
│   └── utils/               # 公共工具
├── configs/
│   └── config.yaml          # 配置文件
├── go.mod
├── go.sum
└── Makefile