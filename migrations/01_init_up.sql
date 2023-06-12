CREATE TABLE IF NOT EXISTS users
(
    id            bigserial primary key,
    login         varchar(255)   not null,
    password_hash varchar(255)   not null,
    CONSTRAINT    login_unique UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS versions
(
    version     bigint not null,
    user_id     bigint not null,
    CONSTRAINT  uniq_version_user_id UNIQUE (version, user_id),
    CONSTRAINT  versions_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE IF NOT EXISTS cards
(
    id          bigserial primary key,
    user_id     bigint not null,
    number      varchar(255),
    holder      varchar(255),
    expire      varchar(255),
    cvc         varchar(255),
    metadata    varchar(255),
    CONSTRAINT  cards_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

create table if not exists accounts
(
    id          bigserial primary key,
    user_id     bigint not null,
    login       varchar(255),
    password    varchar(255),
    metadata     varchar(255),
    CONSTRAINT  accounts_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

create table if not exists notes
(
    id          bigserial primary key,
    user_id     bigint not null,
    text        varchar(255),
    metadata     varchar(255),
    CONSTRAINT  notes_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);

create table if not exists binary
(
    id          bigserial primary key,
    user_id     bigint not null,
    body        bytea,
    metadata     varchar(255),
    CONSTRAINT binary_users_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);
