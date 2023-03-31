drop table if exists users;
drop table if exists roles;
drop table if exists persons;
drop table if exists roles_persons;
drop table if exists content;
drop table if exists films;
drop table if exists films_persons;
drop table if exists countries;
drop table if exists genres;
drop table if exists content_countries;
drop table if exists content_genres;
drop table if exists series;
drop table if exists episodes;

CREATE SCHEMA filmium;
SET search_path=filmium;

CREATE DOMAIN gender CHAR(1)
    CHECK (value IN ( 'F' , 'M' ));

create table users (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    date_birth date not null,
    avatar_url text
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

create table roles_persons (
    role_id bigint references roles(id) on delete cascade,
    person_id bigint references persons(id) on delete cascade,
    PRIMARY KEY (role_id, person_id)
);

create table content (
    id bigserial primary key,
    title text,
    description text,
    rating numeric(2, 2),
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

create table films_persons (
    film_id bigint references films(id) on delete cascade,
    person_id bigint references persons(id) on delete cascade,
    primary key (film_id, person_id)
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
