CREATE TABLE IF NOT EXISTS butter.users (
    id varchar(40) not null,
    username varchar(50) not null unique,
    password varchar(100) not null,
    name varchar(50) not null,
    email varchar(50) not null unique,
    birthdate timestamp null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    deleted_at timestamp null,
    primary key (id)
) engine = InnoDB;