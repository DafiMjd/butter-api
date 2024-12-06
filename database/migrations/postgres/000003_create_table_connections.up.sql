
CREATE TABLE IF NOT EXISTS butter.connections (
    followee_id UUID REFERENCES butter.users(id),
    followee_username VARCHAR(50) REFERENCES butter.users(username),
    follower_id UUID REFERENCES butter.users(id),
    follower_username VARCHAR(50) REFERENCES butter.users(username),
    PRIMARY KEY (followee_id, follower_id),
    CONSTRAINT fk_followee_id FOREIGN KEY (followee_id) REFERENCES butter.users(id),
    CONSTRAINT fk_follower_id FOREIGN KEY (follower_id) REFERENCES butter.users(id),
    CONSTRAINT fk_followee_username FOREIGN KEY (followee_username) REFERENCES butter.users(username),
    CONSTRAINT fk_follower_username FOREIGN KEY (follower_username) REFERENCES butter.users(username)
);