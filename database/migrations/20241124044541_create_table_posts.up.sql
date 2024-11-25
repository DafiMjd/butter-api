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