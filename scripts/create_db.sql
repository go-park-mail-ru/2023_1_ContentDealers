drop table if exists users cascade;
drop table if exists roles cascade;
drop table if exists persons cascade;
drop table if exists content_roles_persons cascade;
drop table if exists content cascade;
drop table if exists films cascade;
drop table if exists countries cascade;
drop table if exists genres cascade;
drop table if exists content_countries cascade;
drop table if exists content_genres cascade;
drop table if exists series cascade;
drop table if exists episodes cascade;
drop table if exists selections cascade;
drop table if exists content_selections cascade;
drop table if exists users_content_favorites cascade;
drop table if exists users_persons_favorites cascade;

-- namespace, gender, function set_timestamp

create schema if not exists filmium;
create schema if not exists user_schema;
create schema if not exists favorites_schema;
set search_path=filmium;

create extension if not exists pg_trgm;

drop domain if exists gender cascade;
create domain gender char(1)
    check (value IN ('F', 'M'));

drop type if exists filmium.content_type cascade;
create type content_type as enum (
    'film',
    'series'
);

create or replace function set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

-- tables

create table user_schema.users (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    avatar_url text not null default 'media/avatars/default_avatar.jpg',
    sub_expiration date not null default date('1970-01-01'),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create table favorites_schema.users_content_favorites (
    user_id bigint not null,
    content_id bigint not null,
    created_at timestamp not null default now(),
    primary key (user_id, content_id)
);

create table favorites_schema.users_persons_favorites (
    user_id bigint not null,
    person_id bigint not null,
    created_at timestamp not null default now(),
    primary key (user_id, person_id)
);

create table roles (
    id bigserial primary key,
    title text unique not null
);

create table persons (
    id bigserial primary key,
    name text not null,
    gender filmium.gender not null,
    growth integer,
    birthplace text,
    avatar_url text not null default 'media/avatars/default_avatar.jpg',
    age integer
);

create table content (
    id bigserial primary key,
    title text not null,
    description text not null,
    rating numeric(4, 2),
    year integer,
    is_free boolean not null default true,
    age_limit integer not null default 0,
    preview_url text not null,
    trailer_url text not null,
    type content_type not null
);

create table films (
    id bigserial primary key,
    content_id bigint not null references content(id) on delete cascade,
    content_url text not null
);

create table selections (
    id bigserial primary key,
    title text not null
);

create table content_selections (
    content_id bigint references content(id) on delete cascade,
    selection_id bigint references selections(id) on delete cascade,
    primary key (content_id, selection_id)
);

create table content_roles_persons (
    role_id bigint references roles(id) on delete cascade,
    person_id bigint references persons(id) on delete cascade,
    content_id bigint references content(id) on delete cascade,
    primary key (role_id, person_id, content_id)
);

create table countries (
    id bigserial primary key,
    name text not null
);

create table genres (
    id bigserial primary key,
    name text not null
);

create table content_countries (
    content_id bigint references content(id) on delete cascade,
    country_id bigint references countries(id) on delete cascade,
    primary key (content_id, country_id)
);

create table content_genres (
    content_id bigint references content(id) on delete cascade,
    genre_id bigint references genres(id) on delete cascade,
    primary key (content_id, genre_id)
);

create table series (
    id bigserial primary key,
    content_id bigint not null references content(id) on delete cascade
);

create table episodes (
    id bigserial primary key,
    series_id bigint not null references series(id) on delete cascade,
    season_num integer not null,
    episode_num integer not null,
    content_url text not null,
    release_date date,
    title text 
);

-- trigger
create trigger set_timestamp_users
before update on user_schema.users
for each row
execute procedure set_timestamp();

