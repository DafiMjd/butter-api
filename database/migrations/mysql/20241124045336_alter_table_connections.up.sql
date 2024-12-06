ALTER TABLE butter.connections ADD followee_username VARCHAR(50) AFTER followee_id;

ALTER TABLE butter.connections ADD follower_username VARCHAR(50) AFTER follower_id;
