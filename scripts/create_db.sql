drop schema if exists user_schema cascade;
drop schema if exists content_schema cascade;
drop schema if exists user_action_schema cascade;

-- namespace, gender

create schema if not exists content_schema;
create schema if not exists user_schema;
create schema if not exists user_action_schema;

create extension if not exists pg_trgm;

set search_path = 'content_schema';

drop domain if exists gender cascade;
create domain gender char(1)
    check (value IN ('F', 'M'));

drop type if exists content_schema.content_type cascade;
create type content_type as enum (
    'film',
    'series'
);

-- tables

-- 3 нормальная форма
create table user_schema.users (
    id bigserial primary key,
    email text not null unique, -- пользователя идентифицируем по email
    password_hash text not null,
    avatar_url text not null default 'media/avatars/default_avatar.jpg',
    sub_expiration date not null default date('1970-01-01'),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- 3 нормальная форма
create table user_action_schema.users_content_favorites (
    user_id bigint not null,
    content_id bigint not null,
    created_at timestamp not null default now(),
    primary key (user_id, content_id)
);

-- 3 нормальная форма
create table user_action_schema.users_persons_favorites (
    user_id bigint not null,
    person_id bigint not null,
    created_at timestamp not null default now(),
    primary key (user_id, person_id)
);

-- 3 нормальная форма
create table user_action_schema.ratings (
    user_id bigint not null,
    content_id bigint not null,
    rating numeric(4, 2),
    created_at timestamp not null default now(),
    primary key (user_id, content_id)
);

-- 3 нормальная форма
create table user_action_schema.history_views (
    user_id bigint not null,
    content_id bigint not null,
    stop_view bigint not null,
    duration bigint not null,
    created_at timestamp not null default now(),
    primary key (user_id, content_id)
);

-- 3 нормальная форма
create table content_schema.roles (
    id bigserial primary key,
    title text unique not null -- Одна запись - одна роль
);

-- 3 нормальная форма
create table content_schema.persons (
    id bigserial primary key,
    name text not null,
    gender content_schema.gender not null,
    growth integer,
    birthplace text,
    avatar_url text not null default 'media/avatars/default_avatar.jpg',
    age integer
);

-- 3 нормальная форма
create table content_schema.content (
    id bigserial primary key,
    title text not null,
    description text not null,
    rating numeric(3, 1),
    year integer,
    is_free boolean not null default true,
    age_limit integer not null default 0,
    preview_url text not null,
    trailer_url text not null,
    type content_type not null,

    sum_ratings numeric(12, 2) not null default 0,
    count_ratings bigint not null default 0
);

-- 3 нормальная форма
create table content_schema.films (
    id bigserial primary key,
    content_id bigint not null references content(id) on delete cascade,
    content_url text not null
);

-- 3 нормальная форма
create table content_schema.selections (
    id bigserial primary key,
    title text not null
);

-- 3 нормальная форма
create table content_schema.content_selections (
    content_id bigint references content(id) on delete cascade,
    selection_id bigint references selections(id) on delete cascade,
    primary key (content_id, selection_id)
);

-- 3 нормальная форма
create table content_schema.content_roles_persons (
    role_id bigint references roles(id) on delete cascade,
    person_id bigint references persons(id) on delete cascade,
    content_id bigint references content(id) on delete cascade,
    primary key (role_id, person_id, content_id)
);

-- 3 нормальная форма
create table content_schema.countries (
    id bigserial primary key,
    name text not null
);

-- 3 нормальная форма
create table content_schema.genres (
    id bigserial primary key,
    name text not null
);

-- 3 нормальная форма
create table content_schema.content_countries (
    content_id bigint references content(id) on delete cascade,
    country_id bigint references countries(id) on delete cascade,
    primary key (content_id, country_id)
);

-- 3 нормальная форма
create table content_schema.content_genres (
    content_id bigint references content(id) on delete cascade,
    genre_id bigint references genres(id) on delete cascade,
    primary key (content_id, genre_id)
);

-- 3 нормальная форма
create table content_schema.series (
    id bigserial primary key,
    content_id bigint not null references content(id) on delete cascade
);

-- 3 нормальная форма
create table content_schema.episodes (
    id bigserial primary key,
    series_id bigint not null references series(id) on delete cascade,
    preview_url text default 'previews_episodes/default_preview.jpg',
    season_num integer not null,
    episode_num integer not null,
    content_url text not null,
    release_date date,
    title text 
);

-- indexes

-- foreign key indexes
create index if not exists content_countries__country_id on content_schema.content_countries(country_id);
create index if not exists content_genres__genre_id on content_schema.content_genres(genre_id);
create index if not exists content_roles_persons__content_id on content_schema.content_roles_persons(content_id);
create index if not exists content_roles_persons__person_id on content_schema.content_roles_persons(person_id);
create index if not exists content_selections__selection_id on content_schema.content_selections(selection_id);
create index if not exists episodes__series_id on content_schema.episodes(series_id);
create index if not exists films__content_id on content_schema.films(content_id);
create index if not exists series__content_id on content_schema.series(content_id);

-- search indexes
create index if not exists content__title on content_schema.content using gin (title public.gin_trgm_ops);
create index if not exists persons__name on content_schema.persons using gin (name public.gin_trgm_ops);

-- triggers and functions

create or replace function user_schema.set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

create trigger set_timestamp_users
before update on user_schema.users
for each row
execute procedure user_schema.set_timestamp();



create or replace function content_schema.update_rating()
returns trigger as $$
begin
    new.rating = new.sum_ratings / new.count_ratings;
    return new;
end;
$$ language plpgsql;

create trigger update_rating_trigger
before update of sum_ratings, count_ratings on content_schema.content
for each row
execute function content_schema.update_rating();
