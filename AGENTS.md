# AGENTS.md

## Stack

- Go 1.25 backend, module `github.com/MelodicTechno/anime-list`
- Gin web framework, Gorm ORM, PostgreSQL, Redis cache
- RESTful API style

## Project structure

```
cmd/api/main.go         → entry point
internal/
  config/               → configuration management
  database/             → DB/Redis connection setup (infrastructure)
  handler/              → Gin HTTP handlers
  service/              → business logic
  repository/           → Gorm data access
  model/                → shared data structs
pkg/utils/              → public utilities
configs/config.yaml     → YAML config file
```

**Layered flow:** `handler → service → repository → PostgreSQL (+ Redis cache)`

## Model conventions

- IDs use `int64`, timestamps use `time.Time`
- JSON tags: camelCase (`animeId`, `createdAt`)
- No explicit column tags — Gorm auto-infers snake_case from field names (`AnimeID` → `anime_id`)
- All models in `internal/model/` with `package model`

## Model conventions

- IDs use `int64`, timestamps use `time.Time`
- JSON tags: camelCase (`animeId`, `createdAt`)
- DB column tags: snake_case (`anime_id`, `created_at`)
- All models in `internal/model/` with `package model`

## Build commands

```
make build             → go build -o bin/api.exe ./cmd/api/
make run               → go run ./cmd/api/
./scripts/build.ps1    → PowerShell build script
bash scripts/build.sh  → bash build script
```

## Current state

Early-stage project. Only model structs exist (`internal/model/anime.go`, `comment.go`). All other layers are empty directories. No build/test/lint commands are configured yet (Makefile is empty, no CI).

## Existing instruction files

- `.trae/rules.md` — original assistant rules (Chinese, same content as this file's intent)
