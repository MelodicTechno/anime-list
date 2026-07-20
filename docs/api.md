# API 接口

Base URL: `http://localhost:8080`

## 用户

| 方法 | 路径 | 说明 |
|---|---|---|
| POST | `/api/register` | 注册 |
| POST | `/api/login` | 登录 |

### POST /api/register

```
{"username": "foo", "email": "foo@example.com", "password": "123456"}
```

### POST /api/login

```
{"account": "foo", "password": "123456"}
```
