# API 接口

Base URL: `http://localhost:8080`

## 用户

| 方法 | 路径 | 认证 | 说明 |
|---|---|---|---|
| POST | `/api/register` | 无 | 注册 |
| POST | `/api/login` | 无 | 登录，返回双 token |
| POST | `/api/refresh` | 无 | 刷新 access_token |
| GET | `/api/me` | Bearer access_token | 获取当前用户 ID |

### POST /api/register

```json
{"username": "foo", "email": "foo@example.com", "password": "123456"}
```

### POST /api/login

```json
{"account": "foo", "password": "123456"}
```
响应：
```json
{
  "data": {
    "accessToken": "...",
    "refreshToken": "...",
    "user": {"id": 1, "username": "foo", ...}
  }
}
```

### POST /api/refresh

```json
{"refreshToken": "..."}
```

### GET /api/me

Header: `Authorization: Bearer <accessToken>`

## 书架

所有接口需要 Bearer access_token。

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/api/bookshelves` | 创建书架 |
| GET | `/api/bookshelves` | 获取用户的所有书架 |
| GET | `/api/bookshelves/:id` | 获取单个书架（含作品列表） |
| PUT | `/api/bookshelves/:id` | 修改书架名称 |
| DELETE | `/api/bookshelves/:id` | 删除书架 |
| POST | `/api/bookshelves/:id/items` | 添加作品到书架 |
| DELETE | `/api/bookshelves/:id/items/:itemId` | 从书架移除作品 |

### POST /api/bookshelves

```json
{"name": "正在追"}
```

### POST /api/bookshelves/:id/items

```json
{"animeId": 1, "stateId": 2}
```
`stateId` 可选：1 想看、2 在看、3 已看、4 弃了。
