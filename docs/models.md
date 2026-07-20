# 数据模型

## Anime（动画/漫画）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| Title | `string` | `title` | `title` | 标题 |
| ReleaseDate | `time.Time` | `releaseDate` | `release_date` | 发布日期 |
| Score | `float64` | `score` | `score` | 综合评分（0-10） |

## User（用户）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| Username | `string` | `username` | `username` | 用户名 |
| Email | `string` | `email` | `email` | 邮箱 |
| PasswordHash | `string` | `-` | `password_hash` | 密码哈希（JSON 隐藏） |
| Avatar | `string` | `avatar` | `avatar` | 头像 URL |
| CreatedAt | `time.Time` | `createdAt` | `created_at` | 注册时间 |
| UpdatedAt | `time.Time` | `updatedAt` | `updated_at` | 更新时间 |

## Comment（评论）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| AnimeID | `int64` | `animeId` | `anime_id` | 关联动画 ID |
| UserID | `int64` | `userId` | `user_id` | 用户 ID |
| Content | `string` | `content` | `content` | 评论内容 |
| CreatedAt | `time.Time` | `createdAt` | `created_at` | 创建时间 |
| UpdatedAt | `time.Time` | `updatedAt` | `updated_at` | 更新时间 |

## Bookshelf（书架）

用户自定义的作品分类，用于整理和管理已看/想看的动漫漫画。

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| UserID | `int64` | `userId` | `user_id` | 所属用户 ID |
| Name | `string` | `name` | `name` | 书架名称 |
| CreatedAt | `time.Time` | `createdAt` | `created_at` | 创建时间 |

### BookshelfItem（书架条目）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| BookshelfID | `int64` | `bookshelfId` | `bookshelf_id` | 书架 ID |
| AnimeID | `int64` | `animeId` | `anime_id` | 作品 ID |

> Bookshelf ↔ Anime 为多对多关系，通过 BookshelfItem 关联。

## Category（类型/标签）

作品的分类标签（如"热血"、"恋爱"、"科幻"等）。

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| Name | `string` | `name` | `name` | 分类名称 |

### AnimeCategory（动漫-类型关联）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| AnimeID | `int64` | `animeId` | `anime_id` | 作品 ID |
| CategoryID | `int64` | `categoryId` | `category_id` | 类型 ID |

> Anime ↔ Category 为多对多关系，通过 AnimeCategory 关联。

## Favorite（收藏夹）

用户建立的收藏夹。

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| UserID | `int64` | `userId` | `user_id` | 所属用户 ID |
| Name | `string` | `name` | `name` | 收藏夹名称 |
| CreatedAt | `time.Time` | `createdAt` | `created_at` | 创建时间 |

### FavoriteItem（收藏条目）

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| FavoriteID | `int64` | `favoriteId` | `favorite_id` | 收藏夹 ID |
| AnimeID | `int64` | `animeId` | `anime_id` | 作品 ID |

> Favorite ↔ Anime 为多对多关系，通过 FavoriteItem 关联。

## WatchPlan（计划表）

记录用户计划观看的作品及观看状态。

| 字段 | 类型 | JSON | 数据库列 | 说明 |
|---|---|---|---|---|
| ID | `int64` | `id` | `id` | 主键 |
| UserID | `int64` | `userId` | `user_id` | 用户 ID |
| AnimeID | `int64` | `animeId` | `anime_id` | 作品 ID |
| Status | `string` | `status` | `status` | 状态（planned/watching/completed/dropped） |
| Notes | `string` | `notes` | `notes` | 备注 |
| CreatedAt | `time.Time` | `createdAt` | `created_at` | 创建时间 |
| UpdatedAt | `time.Time` | `updatedAt` | `updated_at` | 更新时间 |

> User ↔ Anime 为一对多关系，每个用户对每部作品有一条计划记录。

## 完整性约定

- 所有 ID 类型：`int64`
- 所有时间戳类型：`time.Time`
- JSON 序列化：camelCase（`animeId`、`createdAt`）
- 数据库列名：Gorm 自动从字段名推断为 snake_case（`AnimeID` → `anime_id`），无需显式标签
- 模型文件位于 `internal/model/`，包名为 `model`

## 实体关系

```
User (1) ──── (N) Comment
User (1) ──── (N) Bookshelf ──── (N) BookshelfItem ──── (1) Anime
User (1) ──── (N) Favorite  ──── (N) FavoriteItem  ──── (1) Anime
User (1) ──── (N) WatchPlan ──── (1) Anime
Anime (1) ──── (N) Comment
Anime (1) ──── (N) AnimeCategory ──── (1) Category
```
