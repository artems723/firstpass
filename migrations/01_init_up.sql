CREATE TABLE IF NOT EXISTS users
(
    id            bigserial primary key,
    login         varchar(255)   not null,
    password_hash varchar(255)   not null,
    CONSTRAINT    login_unique UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS cards
(
    id          bigserial primary key,
    name        varchar(255),
    user_id     varchar(255),
    number      varchar(255),
    holder      varchar(255),
    expire      varchar(255),
    cvc         varchar(255),
        comment    varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    constraint  cards_users_id_fkey foreign key (user_id) references users
);

create table if not exists accounts
(
    id          bigserial primary key,
    name        varchar(255),
    user_id     varchar(255),
    login       varchar(255),
    password    varchar(255),
    comment     varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    constraint  accounts_users_id_fkey foreign key (user_id) references users
);

create table if not exists notes
(
    id          bigserial primary key,
    name        varchar(255),
    user_id     varchar(255),
    text        varchar(255),
    comment     varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    constraint  notes_users_id_fkey foreign key (user_id) references users
);

create table if not exists binary
(
    id          bigserial primary key,
    name        varchar(255),
    user_id     varchar(255),
    body        bytea,
    comment     varchar(255),
    created_at  timestamp,
    updated_at  timestamp,
    constraint notes_users_id_fkey foreign key (user_id) references users
);
