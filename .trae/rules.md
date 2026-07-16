# Anime List

Anime List 是一个用来让我（也支持其他人）标记自己看过的动漫或者漫画，然后打分并记录观看体验和感受的项目

这里是后端项目

## 技术

后端语言使用Go，框架是Gin，使用Gorm连接数据库。也会使用Redis来当作缓存，加快对经常访问的动画或者漫画的查询等。数据库使用的是Postgres

使用RESTful api风格

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

## 特别

尽可能发挥Go并发的优势并利用Redis的性能

## 运行

Trae，你千万不要自己运行什么命令，让我运行就好。