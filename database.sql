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

CREATE TABLE IF NOT EXISTS butter.connection (
    followee_id varchar(40) not null,
    follower_id varchar(40) not null,
    primary key (followee_id, follower_id),
    foreign key (followee_id) references users(id),
    foreign key (follower_id) references users(id)
) engine = InnoDB;

CREATE TABLE IF NOT EXISTS butter.posts (
    id varchar(40) not null,
    user_id varchar(40) not null,
    content text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    deleted_at timestamp null,
    primary key (id),
    foreign key (user_id) references users(id)
) engine = InnoDB;

