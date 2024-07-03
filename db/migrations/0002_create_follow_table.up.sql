CREATE TABLE follows (
    follower_id uuid NOT NULL,
    followed_id uuid NOT NULL,
    follow_date DATE NOT NULL,
    PRIMARY KEY (follower_id, followed_id),
    FOREIGN KEY (follower_id) REFERENCES users(user_id),
    FOREIGN KEY (followed_id) REFERENCES users(user_id)
);