drop table if exists users;
drop table if exists movie_selections;
drop table if exists movies;
drop table if exists selections;

-- drop sequence if exists items_id_seq;
-- create sequence items_id_seq increment 1 minvalue 1 maxvalue 2147483647 start 1 cache 1;

create or replace function trigger_set_timestamp()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

create table users (
    id bigserial primary key,
    email text not null unique,
    password_hash text not null,
    created_at timestamp not null default now(),
    -- timestamptz
    updated_at timestamp not null default now()
);


create table movies (
    id bigserial primary key,
    title text not null unique,
    description text,
    preview_url text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create table selections (
    id bigserial primary key,
    title text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create table movie_selections (
    id bigserial primary key,
    movie_id int,
    selection_id int,
    foreign key(movie_id) references movies(id) on delete cascade,
    foreign key(selection_id) references selections(id) on delete cascade,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create trigger set_timestamp_users
before update on users
for each row
execute procedure trigger_set_timestamp();

create trigger set_timestamp_movies
before update on movies
for each row
execute procedure trigger_set_timestamp();

create trigger set_timestamp_selections
before update on selections
for each row
execute procedure trigger_set_timestamp();

create trigger set_timestamp_movie_selections
before update on movie_selections
for each row
execute procedure trigger_set_timestamp();

insert into movies (id, preview_url, title, description) values 
	(1, 'media/previews/mad-max.jpg', 'Mad Max',  'Mad Max описание фильма'),
	(2, 'media/previews/back-to-the-future.jpg',  'Back to the future',  'Back to the futur описание фильма'),
	(3, 'media/previews/king-kong.jpg',  'King Kong',  'King Kong описание фильма'),
	(4, 'media/previews/terminator.jpg',  'Terminator',  'Terminator описание фильма'),
	(5, 'media/previews/godzilla.jpg',  'Godzilla',  'Godzilla описание фильма'),
	(6, 'media/previews/007.jpg',  '007',  '007 описание фильма'),
	(7, 'media/previews/black-panther.jpg',  'Back Panther',  'Back Panther описание фильма'),
	(8, 'media/previews/captain-america.jpg',  'Capitan America',  'Capitan America описани фильма'),
	(9, 'media/previews/pacific-rim.jpg',  'Pacific Rim',  'Pacific Rim описание фильма'),
	(10, 'media/previews/interstellar.jpg',  'Interstellar',  'Interstellar описание фильма'),
	(11, 'media/previews/face.jpg',  'Face',  'Face описание фильма'),
	(12, 'media/previews/thor.jpg',  'Thor',  'Thor описание фильма'),
	(13, 'media/previews/dune.jpg',  'Dune',  'Dune описание фильма'),
	(14, 'media/previews/avatar.jpg',  'Avatar',  'Avatar описание фильма'),
	(15, 'media/previews/star-wars.jpg',  'Star Wars',  'Star Wars описание фильма'),
	(16, 'media/previews/venom.jpg',  'Vanom',  'Vanom описание фильма');

insert into selections (id, title) values 
    (1, 'Filmium рекомендует'),
    (2, 'Лучшие боевики'),
    (3, 'Фэнтези');

insert into movie_selections (movie_id, selection_id) values 
    (2, 1),
    (3, 1),
    (6, 1),
    (7, 1),
    (8, 1),
    (14, 1),

    (1, 2),
    (4, 2),
    (5, 2),
    (6, 2),
    (9, 2),

    (10, 3),
    (13, 3),
    (14, 3),
    (15, 3);

