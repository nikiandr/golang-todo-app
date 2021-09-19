CREATE TABLE users
(
    id            serial primary key,
    name          varchar(255) not null,
    surname       varchar(255),
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial primary key,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial primary key,
    list_id     int references todo_lists (id) on delete cascade not null,
    title       varchar(255)                                     not null,
    description varchar(255),
    done        boolean                                          not null default false
);
