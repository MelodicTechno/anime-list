-- 使用说明:
--   psql -U postgres -f cmd/sql/create_database.sql   (首次运行，需先连接 postgres 库)
--   psql -U postgres -d anime_list -f cmd/sql/create_database.sql  (已有 anime_list 库时重建表)

drop database if exists anime_list;
create database anime_list;

\c anime_list

drop table if exists watch_plans      cascade;
drop table if exists anime_categories cascade;
drop table if exists favorite_items   cascade;
drop table if exists bookshelf_items  cascade;
drop table if exists comments         cascade;
drop table if exists favorites        cascade;
drop table if exists bookshelves      cascade;
drop table if exists categories       cascade;
drop table if exists anime            cascade;
drop table if exists users            cascade;

create table users (
    id           bigserial primary key,
    username     varchar(64)  not null unique,
    email        varchar(128) not null unique,
    password_hash varchar(256) not null,
    avatar       varchar(512) not null default '',
    created_at   timestamp    not null default now(),
    updated_at   timestamp    not null default now()
);

create table anime (
    id           bigserial primary key,
    title        varchar(256) not null,
    release_date date,
    score        numeric(3,1) not null default 0,
    constraint chk_score_range check (score >= 0 and score <= 10)
);

create table categories (
    id   bigserial primary key,
    name varchar(64) not null unique
);

create table comments (
    id         bigserial primary key,
    anime_id   bigint    not null references anime(id) on delete cascade,
    user_id    bigint    not null references users(id) on delete cascade,
    content    text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
create index idx_comments_anime_id on comments(anime_id);
create index idx_comments_user_id  on comments(user_id);

create table bookshelves (
    id           bigserial primary key,
    user_id      bigint      not null references users(id) on delete cascade,
    name         varchar(64) not null,
    created_at   timestamp   not null default now(),
    unique(user_id, name)
);

create table bookshelf_items (
    id           bigserial primary key,
    bookshelf_id bigint not null references bookshelves(id) on delete cascade,
    anime_id     bigint not null references anime(id) on delete cascade,
    unique(bookshelf_id, anime_id)
);
create index idx_bookshelf_items_anime on bookshelf_items(anime_id);

create table favorites (
    id         bigserial primary key,
    user_id    bigint      not null references users(id) on delete cascade,
    name       varchar(64) not null,
    created_at timestamp   not null default now(),
    unique(user_id, name)
);

create table favorite_items (
    id          bigserial primary key,
    favorite_id bigint not null references favorites(id) on delete cascade,
    anime_id    bigint not null references anime(id) on delete cascade,
    unique(favorite_id, anime_id)
);
create index idx_favorite_items_anime on favorite_items(anime_id);

create table anime_categories (
    anime_id    bigint not null references anime(id) on delete cascade,
    category_id bigint not null references categories(id) on delete cascade,
    primary key (anime_id, category_id)
);
create index idx_anime_categories_category on anime_categories(category_id);

create table watch_plans (
    id         bigserial primary key,
    user_id    bigint    not null references users(id) on delete cascade,
    anime_id   bigint    not null references anime(id) on delete cascade,
    status     varchar(16) not null default 'planned',
    notes      text      not null default '',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    unique(user_id, anime_id),
    constraint chk_watch_plan_status check (status in ('planned', 'watching', 'completed', 'dropped'))
);
create index idx_watch_plans_anime on watch_plans(anime_id);
