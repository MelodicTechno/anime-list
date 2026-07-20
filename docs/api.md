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
