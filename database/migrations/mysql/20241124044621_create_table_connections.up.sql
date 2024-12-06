
CREATE TABLE IF NOT EXISTS butter.connections (
    followee_id varchar(40) not null,
    follower_id varchar(40) not null,
    primary key (followee_id, follower_id),
    foreign key (followee_id) references users(id),
    foreign key (follower_id) references users(id)
) engine = InnoDB;