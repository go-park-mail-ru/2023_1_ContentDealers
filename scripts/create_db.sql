drop table if exists users cascade;
drop table if exists roles cascade;
drop table if exists persons cascade;
drop table if exists films_roles_persons cascade;
drop table if exists content cascade;
drop table if exists films cascade;
drop table if exists countries cascade;
drop table if exists genres cascade;
drop table if exists content_countries cascade;
drop table if exists content_genres cascade;
drop table if exists series cascade;
drop table if exists episodes cascade;
drop table if exists selections cascade;
drop table if exists films_selections cascade;

-- namespace, gender, function set_timestamp

CREATE SCHEMA filmium;
SET search_path=filmium;

CREATE DOMAIN gender CHAR(1)
    CHECK (value IN ( 'F' , 'M' ));

create or replace function set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

-- tables

create table users (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    date_birth date not null,
    avatar_url text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create table roles (
    id bigserial primary key,
    title text
);

create table persons (
    id bigserial primary key,
    name text,
    gender filmium.gender,
    growth integer,
    birthplace text,
    avatar_url text,
    age integer
);


create table content (
    id bigserial primary key,
    title text,
    description text,
    rating numeric(4, 2),
    year integer,
    is_free boolean,
    age_limit integer,
    preview_url text
);

create table films (
    id bigserial primary key,
    content_id bigint references content(id) on delete cascade,
    content_url text,
    trailer_url text
);

create table selections (
    id bigserial primary key,
    title text
);

create table films_selections (
    film_id bigint references films(id) on delete cascade,
    selection_id bigint references selections(id) on delete cascade,
    PRIMARY KEY (film_id, selection_id)
);

create table films_roles_persons (
    role_id bigint references roles(id) on delete cascade,
    person_id bigint references persons(id) on delete cascade,
    film_id bigint references films(id) on delete cascade,
    PRIMARY KEY (role_id, person_id, film_id)
);

create table countries (
    id bigserial primary key,
    name text
);

create table genres (
    id bigserial primary key,
    name text
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
    content_id bigint references content(id) on delete cascade
);

create table episodes (
    id bigserial primary key,
    series_id bigint references series(id) on delete cascade,
    season_num integer,
    content_url text,
    title text
);

-- trigger

create trigger set_timestamp_users
before update on users
for each row
execute procedure set_timestamp();
